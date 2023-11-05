package controllers

import (
	"net/http"

	"github.com/cr1m3s/tch_backend/models"
	"github.com/cr1m3s/tch_backend/services"

	"github.com/gin-gonic/gin"
)

type AdvertisementsController struct {
	advertisementService *services.AdvertisementService
	userService          *services.UserService
}

func NewAdvertisementsController() *AdvertisementsController {
	return &AdvertisementsController{
		advertisementService: services.NewAdvertisementService(),
		userService:          services.NewUserService(),
	}
}

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
func (t *AdvertisementsController) AdvCreate(ctx *gin.Context) {
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

	advertisement, err := t.advertisementService.AdvCreate(ctx, inputModel, userID, *t.userService)
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
func (t *AdvertisementsController) AdvPatch(ctx *gin.Context) {
	var inputModel models.AdvertisementUpdate
	if err := ctx.ShouldBindJSON(&inputModel); err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResponseFailed(err.Error()))
		return
	}

	advertisement, err := t.advertisementService.AdvPatch(ctx, inputModel)
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
func (t *AdvertisementsController) AdvDelete(ctx *gin.Context) {
	userID := ctx.GetInt64("user_id")
	if userID == 0 {
		ctx.JSON(http.StatusBadRequest, models.NewResponseFailed("user id error"))
		return
	}

	var inputModel models.AdvertisementDelete

	err := ctx.ShouldBindJSON(&inputModel)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResponseFailed(err.Error()))
		return
	}

	err = t.advertisementService.AdvDelete(ctx, inputModel.AdvID, userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResponseSuccess("advertisement deleted"))
}

// @Advertisement-getall godoc
// @Summary		GET request to get all advertisement
// @Description	endpoint for getting all advertisements
// @Tags		advertisement-getall
// @Param		Authorization header string true "Insert your access token"
// @Produce		json
// @Success		200 {object} map[string]interface{}
// @Router		/protected/advertisement-getall [get]
func (t *AdvertisementsController) AdvGetAll(ctx *gin.Context) {
	advertisements, err := t.advertisementService.AdvGetAll(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResponseFailed(err.Error()))
	}

	ctx.JSON(http.StatusOK, models.NewResponseSuccess(advertisements))
}

// @Advertisement-filter godoc
// @Summary		POST request to get advertisement based on params in filter
// @Description	endpoint for getting specific advertisements
// @Tags		advertisement-filter
// @Param		Authorization header string true "Insert your access token"
// @Param		advertisement-filter body models.AdvertisementFilter true "advertisement filter"
// @Produce		json
// @Success		200 {object} map[string]interface{}
// @Router		/protected/advertisement-filter [post]
func (t *AdvertisementsController) AdvGetFiltered(ctx *gin.Context) {
	var filter models.AdvertisementFilter
	err := ctx.ShouldBindJSON(&filter)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResponseFailed("Filter params not provided."))
		return
	}

	advertisements, err := t.advertisementService.AdvGetFiltered(ctx, filter)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResponseFailed(err.Error()))
	}

	ctx.JSON(http.StatusOK, models.NewResponseSuccess(advertisements))
}
