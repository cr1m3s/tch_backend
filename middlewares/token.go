package middleware

import (
	"net/http"
	"strings"
	"time"

	db "github.com/cr1m3s/tch_backend/queries"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("your-secret-key")

func GenerateToken(user db.User) (string, error) {
	claims := &db.Claims{
		UserID:   user.ID,
		Username: user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Set token expiration time
			IssuedAt:  time.Now().Unix(),                     // Set token issued at time
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

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
			return secretKey, nil
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
