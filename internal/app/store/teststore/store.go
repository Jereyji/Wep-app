package teststore

import (
	"github.com/Jereyji/Wep-app.git/internal/app/model"
	"github.com/Jereyji/Wep-app.git/internal/app/store"
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
		users: make(map[int]*model.User),
	}

	return s.userRepository
}