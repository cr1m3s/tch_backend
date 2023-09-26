package middleware

import (
	"net/http"
	"strings"

	"github.com/cr1m3s/tch_backend/services"
	db "github.com/cr1m3s/tch_backend/queries"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		token, err := jwt.ParseWithClaims(tokenString, &db.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return services.SecretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*db.Claims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		// You can access claims data here
		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
