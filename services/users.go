package services

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/cr1m3s/tch_backend/models"
	"github.com/cr1m3s/tch_backend/queries"
	"github.com/gin-gonic/gin"
)

type ServiceUsers struct {
	db *queries.Queries
}

func NewServiceUsers(db *queries.Queries) *ServiceUsers {
	return &ServiceUsers{
		db: db,
	}
}

func (t *ServiceUsers) LoginUser(ctx *gin.Context, inputModel models.InLogin) (string, error) {

	user, err := t.db.GetUserByEmail(ctx, inputModel.Email)
	if err != nil {
		return "error", err
	}

	cmpPassword := ComparePassword(user.Password, inputModel.Password)
	if cmpPassword != nil {
		err := fmt.Errorf("invalid email or password")
		return "error", err
	}

	token, err := GenerateToken(user)
	if err != nil {
		return "error", err
	}

	return token, nil
}

func (t *ServiceUsers) SignUpUser(ctx *gin.Context, inputModel queries.User) (queries.User, error) {

	_, err := t.db.GetUserByEmail(ctx, inputModel.Email)

	if (err != nil) && (err != sql.ErrNoRows) {
		err = fmt.Errorf("DB search error.")
		return queries.User{}, err
	}
	if err == nil {
		err = fmt.Errorf("User with such email already registred.")
		return queries.User{}, err
	}

	hashedPassword := HashPassword(inputModel.Password)

	args := &queries.CreateUserParams{
		Name:      inputModel.Name,
		Email:     inputModel.Email,
		Password:  hashedPassword,
		Photo:     "default.jpeg",
		Verified:  false,
		Role:      "user",
		UpdatedAt: time.Now(),
	}

	user, err := t.db.CreateUser(ctx, *args)
	if err != nil {
		return queries.User{}, err
	}

	return user, nil
}

func (t *ServiceUsers) GetUserInfo(ctx *gin.Context, userId uint64) (queries.User, error) {
	user, err := t.db.GetUserById(ctx, userId)
	if err != nil {
		return queries.User{}, err
	}

	return user, nil
}
