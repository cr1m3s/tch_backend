package middleware

import (
	"net/http"
	"strings"

	"github.com/cr1m3s/tch_backend/controllers"
	"github.com/cr1m3s/tch_backend/models"
	"github.com/cr1m3s/tch_backend/services"
	"github.com/gin-gonic/gin"
)

func PasswordMiddleware(controller *controllers.UsersController) gin.HandlerFunc {
	return func(c *gin.Context) {
		authString := c.GetHeader("Authorization")
		if authString == "" {
			c.JSON(http.StatusUnauthorized, models.NewResponseFailed("Unauthorized"))
			c.Abort()
			return
		}
		// exppected token:pswd
		pswdString := strings.Split(authString, ":")

		if len(pswdString) != 2 {
			c.JSON(http.StatusUnauthorized, models.NewResponseFailed("Not all info provided for change."))
			c.Abort()
			return
		}

		pswd := pswdString[1]
		userPswd := controller.GetPassword(c)
		err := services.ComparePassword(userPswd, pswd)

		if err != nil {
			c.JSON(http.StatusUnauthorized, models.NewResponseFailed("Unauthorized"))
			c.Abort()
			return
		}

		c.Next()
	}
}
