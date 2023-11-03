package repositories

import (
	"context"

	"github.com/cr1m3s/tch_backend/queries"
)

type AdvertisementsRepositoryInterface interface {
	CreateAdvertisement(ctx context.Context, arg queries.CreateAdvertisementParams) (queries.Advertisement, error)
	UpdateAdvertisement(ctx context.Context, arg queries.UpdateAdvertisementParams) (queries.Advertisement, error)
	GetAdvertisementByID(ctx context.Context, id int64) (queries.Advertisement, error)
	DeleteAdvertisementByID(ctx context.Context, id int64) error
	GetUserById(ctx context.Context, id int64) (queries.User, error)
}

type AdvertisementsRepository struct {
	q *queries.Queries
}

func NewAdvertisementsRepository() *AdvertisementsRepository {
	return &AdvertisementsRepository{
		q: NewAppRepository(),
	}
}

func (t *AdvertisementsRepository) CreateAdvertisement(ctx context.Context, arg queries.CreateAdvertisementParams) (queries.Advertisement, error) {
	return t.q.CreateAdvertisement(ctx, arg)
}

func (t *AdvertisementsRepository) UpdateAdvertisement(ctx context.Context, arg queries.UpdateAdvertisementParams) (queries.Advertisement, error) {
	return t.q.UpdateAdvertisement(ctx, arg)
}

func (t *AdvertisementsRepository) GetAdvertisementByID(ctx context.Context, id int64) (queries.Advertisement, error) {
	return t.q.GetAdvertisementByID(ctx, id)
}

func (t *AdvertisementsRepository) DeleteAdvertisementByID(ctx context.Context, id int64) error {
	return t.q.DeleteAdvertisementByID(ctx, id)
}

func (t *AdvertisementsRepository) GetUserById(ctx context.Context, id int64) (queries.User, error) {
	return t.q.GetUserById(ctx, id)
}
