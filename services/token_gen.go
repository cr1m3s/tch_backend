package services

import (
	"time"

	db "github.com/cr1m3s/tch_backend/queries"
	"github.com/golang-jwt/jwt"
)

var SecretKey = []byte("your-secret-key")

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

	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}