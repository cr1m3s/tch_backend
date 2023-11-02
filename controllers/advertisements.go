package controllers

import (
	"net/http"

	"github.com/cr1m3s/tch_backend/models"

	"github.com/gin-gonic/gin"
)

// @Advertisement-create godoc
// @Summary		POST request to create advertisement
// @Description	endpoint for advertisement creation
// @Tags		advertisement-create
// @Security 	JWT
// @Param		Authorization header string true "Insert your access token"
// @Param		advertisement-create body models.AdvertisementInput true "advertisement information"
// @Produce		json
// @Success		200 {object} map[string]interface{}
// @Router		/protected/advertisement-create [post]
func (t *UsersController) AdvCreate(ctx *gin.Context) {
	userID := ctx.GetInt64("user_id")
	if userID == 0 {
		ctx.JSON(http.StatusBadRequest, models.NewResponseFailed("user id error"))
		return
	}

	var inputModel models.AdvertisementInput
	if err := ctx.ShouldBindJSON(&inputModel); err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResponseFailed(err.Error()))
		return
	}

	advertisement, err := t.userService.AdvCreate(ctx, inputModel, userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResponseSuccess(advertisement))
}

// @Advertisement-patch godoc
// @Summary		PATCH request to update advertisement
// @Description	endpoint for advertisement update
// @Tags		advertisement-patch
// @Param		Authorization header string true "Insert your access token"
// @Param		advertisement-patch body models.AdvertisementUpdate true "advertisement information"
// @Produce		json
// @Success		200 {object} map[string]interface{}
// @Router		/protected/advertisement-patch [patch]
func (t *UsersController) AdvPatch(ctx *gin.Context) {
	var inputModel models.AdvertisementUpdate
	if err := ctx.ShouldBindJSON(&inputModel); err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResponseFailed(err.Error()))
		return
	}

	advertisement, err := t.userService.AdvPatch(ctx, inputModel)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResponseSuccess(advertisement))
}

// @Advertisement-delete godoc
// @Summary		PATCH request to delete advertisement
// @Description	endpoint for advertisement deletion by id
// @Tags		advertisement-delete
// @Param		Authorization header string true "Insert your access token"
// @Param		advertisement-delete body models.Id true "advertisement id"
// @Produce		json
// @Success		200 {object} map[string]interface{}
// @Router		/protected/advertisement-delete [delete]
func (t *UsersController) AdvDelete(ctx *gin.Context) {
	var advID models.Id

	err := ctx.ShouldBindJSON(&advID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResponseFailed(err.Error()))
		return
	}

	userID := ctx.GetInt64("user_id")
	if userID == 0 {
		ctx.JSON(http.StatusBadRequest, models.NewResponseFailed("user id error"))
		return
	}

	err = t.userService.AdvDelete(ctx, advID.Id, userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResponseSuccess("advertisement deleted"))
}
