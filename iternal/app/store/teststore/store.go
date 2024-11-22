package teststore

import (
	"video/iternal/app/model"
	"video/iternal/app/store"
)

type Store struct {
	userRepository *UserRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		user:  make(map[string]*model.User),
	}

	return s.userRepository

}
