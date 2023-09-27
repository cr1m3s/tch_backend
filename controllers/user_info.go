package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//	@User_info		godoc
//	@Summary		Get request to see user info
//	@Description	requires valid token
//	@Tags			user_info
//	@Security		JWT
//	@Param			Authorization	header	string	true	"Insert your access token"
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/protected/userinfo [get]
func (ac *AuthController) GetUserInfo(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(int32)

	user, err := ac.db.GetUserById(ctx, userID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "failed to get user info"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}
