package repositories

import (
	"github.com/cr1m3s/tch_backend/queries"
	"context"
)

type AdvertisementsRepositoryInterface interface {
	CreateAdvertisement(ctx context.Context, arg queries.CreateAdvertisementParams)(queries.Advertisement, error)
}

func (t *UsersRepository) CreateAdvertisement(ctx context.Context, arg queries.CreateAdvertisementParams) (queries.Advertisement, error){
	return t.q.CreateAdvertisement(ctx, arg)
}