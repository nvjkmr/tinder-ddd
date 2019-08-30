package persistance

import (
	"errors"
	"github.com/nvjkmr/tinder-ddd/domain/entities"
)

type UserRepositoryImpl struct {
	DB map[int64]entities.User
}

func nextEntityId() {
	return int64(len(r.DB) + 1)
}

func (r UserRepositoryImpl) Get(id int64) (entities.User, error) {
	if r.DB == nil {
		return nil, errors.New("database error")
	}

	if r.DB[id] == nil {
		return nil, errors.New("user not found")
	}

	return r.DB[id], nil
}

func (r UserRepositoryImpl) GetAll() ([]entities.User, error) {
	if r.DB == nil {
		return nil, errors.New("database error")
	}

	users := []entities.User{}
	for _, user := range r.DB {
		users = append(users, user)
	}

	return users, nil
}

func (r UserRepositoryImpl) Save(user *entities.User) error {
	if user == nil {
		return nil, errors.New("empty user")
	}

	if r.DB == nil {
		return errors.New("database error")
	}

	user.ID = nextEntityId()
	r.DB[user.ID] = user
	return nil
}
