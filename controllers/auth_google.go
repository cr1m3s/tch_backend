package controllers

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cr1m3s/tch_backend/queries"
	"github.com/cr1m3s/tch_backend/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const ProtocolPrefix = "https"
const GoogleCallbackUrl = "/api/auth/login-google-callback"
const GoogleCookieName = "oauthstate"
const GoogleQueryNameState = "state"
const GoogleQueryNameCode = "code"
const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
const RedirectDestinationPage = "/api/"

type AuthGoogleController struct {
	googleOauthConfig *oauth2.Config
	serv              *queries.Queries
}

func NewAuthGoogleController(db *queries.Queries) *AuthGoogleController {

	url := ProtocolPrefix + "://" + os.Getenv("GOOGLE_CALLBACK_DOMAIN") + GoogleCallbackUrl

	return &AuthGoogleController{
		googleOauthConfig: &oauth2.Config{
			RedirectURL:  url,
			ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile",
			},
			Endpoint: google.Endpoint,
		},
		serv: db,
	}
}

func (t *AuthGoogleController) LoginGoogle(ctx *gin.Context) {
	b := make([]byte, 16)
	rand.Read(b)
	googleCookieValue := base64.URLEncoding.EncodeToString(b)
	name := GoogleCookieName
	value := googleCookieValue
	maxAge := 3600
	path := ""
	domain := os.Getenv("GOOGLE_CALLBACK_DOMAIN")
	secure := false
	httpOnly := true
	ctx.SetCookie(name, value, maxAge, path, domain, secure, httpOnly)
	authURL := t.googleOauthConfig.AuthCodeURL(googleCookieValue)
	ctx.Redirect(http.StatusTemporaryRedirect, authURL)
}

type GoogleResponse struct {
	Name  string `json: "name"`
	Email string `json: "email"`
	Data  string `json: "-"`
}

type ServiceUsers struct {
	db *queries.Queries
}

func (t *AuthGoogleController) GetUser(ctx context.Context, userInfo GoogleResponse) (queries.User, error) {
	isEmailExist, err := t.serv.IsUserEmailExist(ctx, userInfo.Email)
	if err != nil {
		fmt.Println("Failed to find email in db")
	}

	var user queries.User

	if isEmailExist {
		user, err = t.serv.GetUserByEmail(ctx, userInfo.Email)
		if err != nil {
			fmt.Println("failed to find user")
		}
	} else {
		fmt.Println("Creating")
		args := queries.CreateUserParams{
			Name:      userInfo.Name,
			Email:     userInfo.Email,
			Password:  userInfo.Email,
			Photo:     "default.jpeg",
			Verified:  false,
			Role:      "user",
			UpdatedAt: time.Now(),
		}
		fmt.Println("Hello")
		user, err = t.serv.CreateUser(ctx, args)
		if err != nil {
			fmt.Println("Faield to create user")
		}
	}
	return user, nil
}

func (t *AuthGoogleController) LoginGoogleInfo(ctx *gin.Context) {
	googleCookieValue, _ := ctx.Cookie(GoogleCookieName)
	googleQueryNameState := ctx.Query(GoogleQueryNameState)
	if googleCookieValue != googleQueryNameState {
		log.Println("WARNING: invalid oauth google state")
	}

	codeStr := ctx.Query(GoogleQueryNameCode)
	token, err := t.googleOauthConfig.Exchange(context.Background(), codeStr)
	if err != nil {
		fmt.Println("code exchange wrong", err.Error())
	}

	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		fmt.Println("failed getting user info", err.Error())
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("failed read user info", err.Error())
	}

	var userInfo GoogleResponse
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		fmt.Println("failed parse response body")
	}
	user, err := t.GetUser(ctx, userInfo)

	if err != nil {
		fmt.Println("Failed to get user by email.")
	}

	tokenJWT, err := services.GenerateToken(user)
	if err != nil {
		fmt.Println("Failed to generate token")
	}

	ctx.JSON(http.StatusTemporaryRedirect, gin.H{"status": "succes", "data": tokenJWT})
	ctx.Redirect(http.StatusTemporaryRedirect, RedirectDestinationPage)
}
