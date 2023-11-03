package di

import (
	"github.com/cr1m3s/tch_backend/repositories"
)

var usersRepository repositories.UsersRepositoryInterface

func NewUsersRepository() repositories.UsersRepositoryInterface {
	if usersRepository == nil {
		usersRepository = repositories.NewUsersRepository()
	}
	return usersRepository
}

func SetUsersRepository(u repositories.UsersRepositoryInterface) {
	usersRepository = u
}

var advertisementsRepository repositories.AdvertisementsRepositoryInterface

func NewAdvertisementsRepository() repositories.AdvertisementsRepositoryInterface {
	if advertisementsRepository == nil {
		advertisementsRepository = repositories.NewAdvertisementsRepository()
	}
	return advertisementsRepository
}

func SetAdvertisementsRepository(u repositories.AdvertisementsRepositoryInterface) {
	advertisementsRepository = u
}
