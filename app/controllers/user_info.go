package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @User_info godoc
// @Summary Get request to see user info
// @Description requires valid token
// @Tags user_info
// @Security JWT
// @Param Authorization header string true "Insert your access token"
// @Produce json
// @Success 200 {object} User
// @Router	/protected/userinfo [get]
func (ac *AuthController) GetUserInfo(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uuid.UUID)

	user, err := ac.db.GetUserById(ctx, userID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "failed to get user info"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}
