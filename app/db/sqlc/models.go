// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package db

import (
	"time"

	"github.com/google/uuid"
	"github.com/golang-jwt/jwt"
)

type User struct {
	ID        uuid.UUID `json:"id"         form:"id"		 `
	Name      string    `json:"name"       form:"name"		 binding:"required"`
	Email     string    `json:"email"      form:"email"		 binding:"required"`
	Photo     string    `json:"photo"      form:"photo"		 `
	Verified  bool      `json:"verified"   form:"verified"   `
	Password  string    `json:"password"   form:"password"	 binding:"required"`
	Role      string    `json:"role"       form:"role"		 `
	CreatedAt time.Time `json:"created_at" form:"created_at" `
	UpdatedAt time.Time `json:"updated_at" form:"updated_at" `
}

type Claims struct {
    UserID   uuid.UUID `json:"user_id"`
    Username string    `json:"username"`
    jwt.StandardClaims
}
