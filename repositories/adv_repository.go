package repositories

import (
	"github.com/cr1m3s/tch_backend/queries"
	"context"
)

type AdvertisementsRepositoryInterface interface {
	CreateAdvertisement(ctx context.Context, arg queries.CreateAdvertisementParams)(queries.Advertisement, error)
	UpdateAdvertisement(ctx context.Context, arg queries.UpdateAdvertisementParams) (queries.Advertisement, error)
	GetAdvertisementByID(ctx context.Context, id int64) (queries.Advertisement, error)
}

func (t *UsersRepository) CreateAdvertisement(ctx context.Context, arg queries.CreateAdvertisementParams) (queries.Advertisement, error){
	return t.q.CreateAdvertisement(ctx, arg)
}

func (t *UsersRepository) UpdateAdvertisement(ctx context.Context, arg queries.UpdateAdvertisementParams) (queries.Advertisement, error) {
	return t.q.UpdateAdvertisement(ctx, arg)
}

func (t *UsersRepository) GetAdvertisementByID(ctx context.Context, id int64) (queries.Advertisement, error){
	return t.q.GetAdvertisementByID(ctx, id)
}