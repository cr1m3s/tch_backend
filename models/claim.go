package models

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type Claims struct {
	UserID   uuid.UUID `json:"user_id"`
	Username string    `json:"username"`
	jwt.StandardClaims
}
