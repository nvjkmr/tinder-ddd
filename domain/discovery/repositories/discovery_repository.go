package repository

import (
	"github.com/nvjkmr/tinder-ddd/domain/discovery/entities"
)

type UserRepository interface {
	Get(ID int64) (entities.User, error)
	GetAll() ([]entities.User, error)
	LikeUser() error
	// Save(user entities.User) error
}

func GetUserRepository() UserRepository {
	return userRepository
}

func InitUserRepository(r UserRepository) {
	userRepository = r
}
