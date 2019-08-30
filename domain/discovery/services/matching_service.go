package service // TODO: figure out if we can have same package name across directories

import (
	"errors"
	"github.com/nvjkmr/tinder-ddd/domain/matcher/entities"
	"github.com/nvjkmr/tinder-ddd/domain/matcher/values"
	"math/rand"
)

type MatchingService interface {
	FindMatch(user entities.User, location values.location) (entities.User, error)
}

type MatchingServiceImpl struct{}

// TODO: find a better way to implement singleton objects
var matchingService MatchingService

// singleton getter
func GetMatchingService() MatchingService {
	return matchingService
}

// singleton setter
func InitMatchingService(m MatchingService) {
	matchingService = m
}

// method
func (s MatchingServiceImpl) FindMatch(user entities.User, location values.Location) (entities.User, error) {
	if user == nil {
		return nil, errors.New("nil user")
	}
	userRepository := repository.GetUserRepository()
	users, err := userRepository.GetAll()

	if err != nil {
		return nil, err
	}

	randomIndex = rand.Intn(len(users))
	match := users[randomIndex]

	// edge case
	if match.ID == user.ID {
		match = users[randomIndex+1]
	}

	return match, nil
}
