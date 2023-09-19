package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/cr1m3s/tch_backend/app/db/sqlc"
	"github.com/cr1m3s/tch_backend/app/utils"
)

type AuthController struct {
	db *db.Queries
}

func NewAuthController(db *db.Queries) *AuthController {
	return &AuthController{db}
}

// @Registraction godoc
// @Summary POST request for registration
// @Description requires username and password for registration
// @Tags register
// @Accept	json
// @Produce json
// @Param email	 query string true   "Email for authentication"
// @Param name query string true "Username for authentication"
// @Param password query string true "Password for authentication"
// @Success 200 {object} map[string]interface{}
// @Router	/api/auth/register [post]
func (ac *AuthController) SignUpUser(ctx *gin.Context) {
	var credentials *db.User

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	hashedPassword := utils.HashPassword(credentials.Password)

	args := &db.CreateUserParams{
		Name:      credentials.Name,
		Email:     credentials.Email,
		Password:  hashedPassword,
		Photo:     "default.jpeg",
		Verified:  false,
		Role:      "user",
		UpdatedAt: time.Now(),
	}
	
	user, err := ac.db.CreateUser(ctx, *args)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"user": user}})
}