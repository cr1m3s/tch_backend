package controllers

import (
	"net/http"

	"github.com/cr1m3s/tch_backend/models"
	"github.com/cr1m3s/tch_backend/services"
	"github.com/gin-gonic/gin"
)

type CategoriesController struct {
	categoriesService *services.CategoriesService
}

func NewCatController() *CategoriesController {
	return &CategoriesController{
		categoriesService: services.NewCategoriesService(),
	}
}

// @Summary		GET all categories parents with children in array
// @Description	endpoint for getting all categories
// @Tags		categories/getall
// @Produce		json
// @Success		200 {object} map[string]interface{}
// @Router		/open/categories/getall [get]
func (t *CategoriesController) CatGetAll(ctx *gin.Context) {
	categories, err := t.categoriesService.CatGetAll(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResponseSuccess(categories))
}
