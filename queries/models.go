// Code generated by sqlc. DO NOT EDIT.
// versions:
// sqlc v1.21.0

package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/golang-jwt/jwt"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Photo     string    `json:"photo"`
	Verified  bool      `json:"verified"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Claims struct {
    UserID   uuid.UUID `json:"user_id"`
    Username string    `json:"username"`
    jwt.StandardClaims
}