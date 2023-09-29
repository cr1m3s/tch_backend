package models

import (
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserID   uint64 `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}
