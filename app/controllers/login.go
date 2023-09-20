package controllers

import (
	"net/http"

	db "github.com/cr1m3s/tch_backend/app/db/sqlc"
	"github.com/cr1m3s/tch_backend/app/utils"
	"github.com/gin-gonic/gin"
)

// @Login godoc
// @Summary POST request for login
// @Description requires username, password and valid JWT token
// @Tags login
// @Accept	json
// @Produce json
// @Param email	 query string true   "Email for authentication"
// @Param password query string true "Password for authentication"
// @Success 200 {object} map[string]interface{}
// @Router	/api/auth/login [post]
func (ac *AuthController) LoginUser(ctx *gin.Context) {
	var user db.User
	email := ctx.Query("email")
	password := ctx.Query("password")

	user, err := ac.db.GetUserByEmail(ctx, email)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	cmpPassword := utils.ComparePassword(user.Password, password)
	
	if cmpPassword != nil{
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}	
 
    if err != nil {
	ctx.JSON(http.StatusBadRequest, gin.H{ "error": gin.H{"info": "Invalid email or password"}})
        return
    }

	ctx.JSON(http.StatusOK, gin.H{"data": "loged in"})
}
