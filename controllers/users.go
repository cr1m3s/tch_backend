package controllers

import (
	"fmt"
	"net/http"

	"github.com/cr1m3s/tch_backend/models"
	"github.com/cr1m3s/tch_backend/queries"
	"github.com/cr1m3s/tch_backend/services"
	"github.com/gin-gonic/gin"
)

type UsersController struct {
	db          *queries.Queries
	userService *services.ServiceUsers
}

func NewUsersController(db *queries.Queries) *UsersController {
	return &UsersController{
		db:          db,
		userService: services.NewServiceUsers(db),
	}
}

// @Login		godoc
// @Summary		POST request for login
// @Description	requires email and password
// @Tags		login
// @Accept		json
// @Produce		json
// @Param		request	body models.InLogin true "request info"
// @Success		200 {object} map[string]interface{}
// @Router		/api/auth/login [post]
func (t *UsersController) LoginUser(ctx *gin.Context) {

	var inputModel models.InLogin
	if err := ctx.ShouldBindJSON(&inputModel); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"data":   err.Error(),
		})
		return
	}

	token, err := t.userService.LoginUser(ctx, inputModel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"data":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "successs",
		"data":   token,
	})
}

// @Registraction godoc
// @Summary		POST request for registration
// @Description	requires username and password for registration
// @Tags		register
// @Accept		json
// @Produce		json
// @Param		user_info body queries.User true "user info for sign in"
// @Success		200	{object} map[string]interface{}
// @Router		/api/auth/register [post]
func (t *UsersController) SignUpUser(ctx *gin.Context) {
	var inputModel queries.User
	if err := ctx.ShouldBindJSON(&inputModel); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"data":   err.Error(),
		})
		return
	}

	user, err := t.userService.SignUpUser(ctx, inputModel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"data":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   user,
	})
}

// @User_info	godoc
// @Summary		Get request to see user info
// @Description	requires valid token
// @Tags		user_info
// @Security	JWT
// @Param		Authorization header string true "Insert your access token"
// @Produce		json
// @Success		200 {object} map[string]interface{}
// @Router		/protected/userinfo [get]
func (t *UsersController) GetUserInfo(ctx *gin.Context) {
	userID := ctx.GetUint64("user_id")
	if userID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"data":   "user id error",
		})
		return
	}

	user, err := t.userService.GetUserInfo(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"data":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   user,
	})

}
