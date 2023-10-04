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

	"github.com/cr1m3s/tch_backend/models"
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
	userService       *services.UserService
}

func NewAuthGoogleController() *AuthGoogleController {

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
		userService: services.NewUserService(),
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

func (t *AuthGoogleController) LoginGoogleCallback(ctx *gin.Context) {
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

	var userInfo models.GoogleResponse
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		fmt.Println("failed parse response body")
	}
	user, err := t.userService.GetOrCreateUser(ctx, userInfo)

	if err != nil {
		fmt.Println("Failed to get user by email.")
	}

	tokenJWT, err := services.GenerateToken(user)
	if err != nil {
		fmt.Println("Failed to generate token")
	}

	ctx.JSON(http.StatusTemporaryRedirect, models.NewResponseSuccess(tokenJWT))
	ctx.Redirect(http.StatusTemporaryRedirect, RedirectDestinationPage)
}
