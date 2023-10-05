package tests

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/cr1m3s/tch_backend/di"
	"github.com/cr1m3s/tch_backend/queries"
	"github.com/cr1m3s/tch_backend/services"
	tests_reposotiries "github.com/cr1m3s/tch_backend/tests/test_repositories"
	"github.com/gin-gonic/gin"
)

var ctx, _ = gin.CreateTestContext(httptest.NewRecorder())

func TestCreateUserDbError(t *testing.T) {
	errExpected := "db search error"
	inputModel := queries.User{}
	want := queries.User{}
	testRepository := tests_reposotiries.NewTestUsersRepository()
	testRepository.ReturnIsUserEmailExist = tests_reposotiries.IsUserEmailExistModel{
		Exist: false,
		Error: errors.New("somthing wrong"),
	}
	di.SetUsersRepository(testRepository)
	service := services.NewUserService()
	got, err := service.UserRegister(ctx, inputModel)

	if err != nil && err.Error() != errExpected {
		t.Errorf("not expected error: [%+v], expected: [%s]", err, errExpected)
	}

	if got != want {
		t.Errorf("not expected user: [%+v], expected: [%+v]", got, want)
	}
}

func TestCreateUserIfEmailAlreadyEsists(t *testing.T) {
	errExpected := "user with such email already registred"
	inputModel := queries.User{}
	want := queries.User{}
	testRepository := tests_reposotiries.NewTestUsersRepository()
	testRepository.ReturnIsUserEmailExist = tests_reposotiries.IsUserEmailExistModel{
		Exist: true,
		Error: nil,
	}
	di.SetUsersRepository(testRepository)
	service := services.NewUserService()
	got, err := service.UserRegister(ctx, inputModel)

	if err != nil && err.Error() != errExpected {
		t.Errorf("not expected error: [%+v], expected: [%s]", err, errExpected)
	}

	if want == got {
		t.Log("OK check expect user model")
	}

	if err == nil {
		t.Error("expected error, but nil returned")
	}

	if want != got {
		t.Errorf("not expected user: [%+v], expected: [%+v]", got, want)
	}
}
