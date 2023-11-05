package controllers

import (
    "net/http"
    "os"

    "github.com/cr1m3s/tch_backend/models"
    "github.com/cr1m3s/tch_backend/services"
    "github.com/gin-gonic/gin"
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/facebook"
)

type AuthFacebookController struct {
	facebookOauthConfig *oauth2.Config
	userService         *services.UserService
}

func GetFacebookOAuthConfig() *oauth2.Config {
    return &oauth2.Config{
        ClientID:     os.Getenv("FACEBOOK_APP_ID"),
        ClientSecret: os.Getenv("FACEBOOK_APP_SECRET"),
        RedirectURL:  os.Getenv("FACEBOOK_OAUTH_REDIRECT_PAGE"),
        Endpoint:     facebook.Endpoint,
        Scopes:       []string{"email"},
    }
}

func GetRandomOAuthStateString() string {
    return "SomeRandomStringAlgorithmForMoreSecurity"
}

func NewAuthFacebookController() *AuthFacebookController {
	return &AuthFacebookController{
		facebookOauthConfig: GetFacebookOAuthConfig(),
		userService: services.NewUserService(),
	}
}

func (afc *AuthFacebookController) LoginFacebook(ctx *gin.Context) {
	var OAuth2Config = GetFacebookOAuthConfig()

    authURL := OAuth2Config.AuthCodeURL(GetRandomOAuthStateString())
    ctx.Redirect(http.StatusTemporaryRedirect, authURL)
}

func (afc *AuthFacebookController) LoginFacebookCallback(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, models.NewResponseSuccess("LoginFacebookCallback function is working!"))
}
