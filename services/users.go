package services

import (
	"fmt"
	"reflect"
	"time"

	"github.com/cr1m3s/tch_backend/configs"
	"github.com/cr1m3s/tch_backend/di"
	"github.com/cr1m3s/tch_backend/models"
	"github.com/cr1m3s/tch_backend/queries"
	"github.com/cr1m3s/tch_backend/repositories"
	"github.com/gin-gonic/gin"
	gomail "gopkg.in/gomail.v2"
)

type UserService struct {
	db repositories.UsersRepositoryInterface
}

func NewUserService() *UserService {
	return &UserService{
		db: di.NewUsersRepository(),
	}
}

func (t *UserService) UserLogin(ctx *gin.Context, inputModel models.InLogin) (string, error) {

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

func (t *UserService) UserRegister(ctx *gin.Context, inputModel queries.User) (queries.User, error) {
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

func (t *UserService) UserInfo(ctx *gin.Context, userId int64) (queries.User, error) {
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

func (t *UserService) UserPatch(ctx *gin.Context, patch queries.User) (queries.User, error) {

	user, err := t.db.GetUserById(ctx, patch.ID)

	userTmp := &queries.UpdateUserParams{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Photo:     user.Photo,
		Verified:  user.Verified,
		Password:  user.Password,
		Role:      user.Role,
		UpdatedAt: user.UpdatedAt,
	}

	if patch.Password != "" {
		patch.Password = HashPassword(patch.Password)
	}

	userValue := reflect.ValueOf(userTmp).Elem()
	patchValue := reflect.ValueOf(patch)
	for i := 0; i < userValue.NumField(); i++ {
		field := userValue.Field(i)
		updateField := patchValue.Field(i)

		if updateField.IsValid() && !updateField.IsZero() {
			field.Set(updateField)
		}
	}

	userTmp.UpdatedAt = time.Now()

	patchedUser, err := t.db.UpdateUser(ctx, *userTmp)
	if err != nil {
		fmt.Println("Faield to update user")
	}

	return patchedUser, nil
}

func (t *UserService) PasswordReset(ctx *gin.Context, email models.EmailRequest) (bool, error) {
	validEmail, _ := t.db.IsUserEmailExist(ctx, email.Email)

	if !validEmail {
		fmt.Println("Email not found")
		return validEmail, fmt.Errorf("Email not found")
	}

	user, err := t.db.GetUserByEmail(ctx, email.Email)
	if err != nil {
		return validEmail, fmt.Errorf("Failed request to DB.")
	}

	_, err = t.EmailSend(email.Email, user)
	if err != nil {
		return false, err
	}

	return validEmail, nil
}

func (t *UserService) EmailSend(userEmail string, user queries.User) (bool, error) {
	token, err := GenerateToken(user)
	if err != nil {
		return false, fmt.Errorf("Failed to generate token.")
	}

	from := configs.GOOGLE_EMAIL_ADDRESS
	response := NewEmail(user.Name, token).message
	msg := gomail.NewMessage()

	msg.SetHeader("From", from)
	msg.SetHeader("To", userEmail)
	msg.SetHeader("Subject", "Password reset")
	msg.SetBody("text/html", response)

	postman := gomail.NewDialer("smtp.gmail.com", 587, from, configs.GOOGLE_EMAIL_SECRET)

	if err := postman.DialAndSend(msg); err != nil {
		return false, fmt.Errorf("Failed to send email.")
	}

	return true, nil
}
