package controllers

import (
	"net/http"

	db "github.com/cr1m3s/tch_backend/queries"
	"github.com/cr1m3s/tch_backend/middlewares"
	"github.com/cr1m3s/tch_backend/services"

	"github.com/gin-gonic/gin"
)

type login_request struct {
	Email     string    `json:"email"      form:"email"		 binding:"required"`
	Password  string    `json:"password"   form:"password"	 binding:"required"`
}

//	@Login			godoc
//	@Summary		POST request for login
//	@Description	requires email and password
//	@Tags			login
//	@Accept			json
//	@Produce		json
//	@Param			request	body		login_request	true	"request info"
//	@Success		200		{object}	map[string]interface{}
//	@Router			/api/auth/login [post]
func (ac *AuthController) LoginUser(ctx *gin.Context) {
	var user db.User
	var req login_request
	
	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "failed", "data": gin.H{"info": "Not enough data provided for log in"}})
		return
	}

	user, err := ac.db.GetUserByEmail(ctx, req.Email)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	cmpPassword := services.ComparePassword(user.Password, req.Password)

	if cmpPassword != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"info": "Invalid email or password"}})
		return
	}
	token, err := middleware.GenerateToken(user)

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
