package teststore

import (
	"video/iternal/app/model"
	"video/iternal/app/store"
)

type UserRepository struct {
	store *Store
	user  map[string]*model.User
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	r.user[u.Email] = u
	u.ID = len(r.user)

	return nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u, ok := r.user[email]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return u, nil
}
