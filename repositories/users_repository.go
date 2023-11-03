package repositories

import (
	"context"

	"github.com/cr1m3s/tch_backend/queries"
)

type UsersRepositoryInterface interface {
	CreateUser(ctx context.Context, arg queries.CreateUserParams) (queries.User, error)
	DeleteUser(ctx context.Context, id int64) error
	GetUserByEmail(ctx context.Context, email string) (queries.User, error)
	GetUserById(ctx context.Context, id int64) (queries.User, error)
	IsUserEmailExist(ctx context.Context, email string) (bool, error)
	ListUsers(ctx context.Context, arg queries.ListUsersParams) ([]queries.User, error)
	UpdateUser(ctx context.Context, arg queries.UpdateUserParams) (queries.User, error)
}

type UsersRepository struct {
	q *queries.Queries
}

func NewUsersRepository() *UsersRepository {
	return &UsersRepository{
		q: NewAppRepository(),
	}
}

func (t *UsersRepository) CreateUser(ctx context.Context, arg queries.CreateUserParams) (queries.User, error) {
	return t.q.CreateUser(ctx, arg)
}

func (t *UsersRepository) DeleteUser(ctx context.Context, id int64) error {
	return t.q.DeleteUser(ctx, id)
}

func (t *UsersRepository) GetUserByEmail(ctx context.Context, email string) (queries.User, error) {
	return t.q.GetUserByEmail(ctx, email)
}

func (t *UsersRepository) GetUserById(ctx context.Context, id int64) (queries.User, error) {
	return t.q.GetUserById(ctx, id)
}

func (t *UsersRepository) IsUserEmailExist(ctx context.Context, email string) (bool, error) {
	return t.q.IsUserEmailExist(ctx, email)
}

func (t *UsersRepository) ListUsers(ctx context.Context, arg queries.ListUsersParams) ([]queries.User, error) {
	return t.q.ListUsers(ctx, arg)
}

func (t *UsersRepository) UpdateUser(ctx context.Context, arg queries.UpdateUserParams) (queries.User, error) {
	return t.q.UpdateUser(ctx, arg)
}
