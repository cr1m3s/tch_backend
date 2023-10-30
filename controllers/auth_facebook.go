package controllers

import (
	"net/http"

	"github.com/cr1m3s/tch_backend/models"
    "github.com/gin-gonic/gin"
)

type AuthFacebookController struct {

}

func NewAuthFacebookController() *AuthFacebookController {
    return &AuthFacebookController{}
}

func (afc *AuthFacebookController) LoginFacebook(c *gin.Context) {
	// Handle Facebook authentication logic here
	c.JSON(http.StatusOK, models.NewResponseSuccess("LoginFacebook function is working!"))
}
