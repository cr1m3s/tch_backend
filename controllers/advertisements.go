package controllers

import (
	"net/http"

	"github.com/cr1m3s/tch_backend/models"
	"github.com/cr1m3s/tch_backend/queries"

	"github.com/gin-gonic/gin"
)

// @Advertisement-create godoc 
// @Summary		POST request to create advertisement
// @Description	endpoint for advertisement creation
// @Tags		advertisement-create 
// @Security 	JWT
// @Param		Authorization header string true "Insert your access token"	 
// @Param		advertisement-create body queries.Advertisement true "advertisement information"
// @Produce		json              
// @Success		200 {object} map[string]interface{}
// @Router		/protected/advertisement-create [post]
func (t *UsersController) AdvCreate(ctx *gin.Context) {
	var inputModel queries.Advertisement
	if err := ctx.ShouldBindJSON(&inputModel); err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResponseFailed(err.Error()))
		return
	}

	advertisement, err := t.userService.AdvCreate(ctx, inputModel)
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
// @Param		advertisement-patch body queries.Advertisement true "advertisement information"
// @Produce		json              
// @Success		200 {object} map[string]interface{}
// @Router		/protected/advertisement-patch [patch]
func (t *UsersController) AdvPatch(ctx *gin.Context) {
	//                            
}                              
                               
// @Advertisement-delete godoc 
// @Summary		PATCH request to delete advertisement
// @Description	endpoint for advertisement deletion by id
// @Tags		advertisement-delete 
// @Param		advertisement-delete body models.Id true "advertisement id"
// @Produce		json              
// @Success		200 {object} map[string]interface{}
// @Router		/protected/advertisement-delete [delete]
func (t *UsersController) AdvDelete(ctx *gin.Context) {
	//                            
}                              