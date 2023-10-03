package services

import (
	"fmt"
	"time"

	"github.com/cr1m3s/tch_backend/models"
	"github.com/cr1m3s/tch_backend/queries"
	"github.com/cr1m3s/tch_backend/repositories"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	db *queries.Queries
}

func NewUserService() *UserService {
	return &UserService{
		db: repositories.NewAppRepository(),
	}
}

func (t *UserService) LoginUser(ctx *gin.Context, inputModel models.InLogin) (string, error) {

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

func (t *UserService) SignUpUser(ctx *gin.Context, inputModel queries.User) (queries.User, error) {
	isEmailExist, err := t.db.IsUserEmailExist(ctx, inputModel.Email)
	if err != nil {
		err = fmt.Errorf("db search error")
		return queries.User{}, err
	}
	if isEmailExist {
		err = fmt.Errorf("user with such email already registred")
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

func (t *UserService) GetUserInfo(ctx *gin.Context, userId int64) (queries.User, error) {
	user, err := t.db.GetUserById(ctx, userId)
	if err != nil {
		return queries.User{}, err
	}

	return user, nil
}

func (t *UserService) GetOrCreateUser(ctx *gin.Context, userInfo models.GoogleResponse) (queries.User, error) {
	isEmailExist, err := t.db.IsUserEmailExist(ctx, userInfo.Email)
	if err != nil {
		fmt.Println("email search query failed")
	}

	var user queries.User

	if isEmailExist {
		user, err = t.db.GetUserByEmail(ctx, userInfo.Email)
		if err != nil {
			fmt.Println("failed to find user")
		}
	} else {
		args := &queries.CreateUserParams{
			Name:      userInfo.Name,
			Email:     userInfo.Email,
			Password:  "",
			Photo:     "default.jpeg",
			Verified:  false,
			Role:      "user",
			UpdatedAt: time.Now(),
		}

		user, err = t.db.CreateUser(ctx, *args)
		if err != nil {
			fmt.Println("Faield to create user")
		}
	}
	return user, nil
}
