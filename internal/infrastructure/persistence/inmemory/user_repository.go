package infrastructure

import (
	"errors"
	domain "example/internal/domain/user"
)

type UserRepository struct {
	users []*domain.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make([]*domain.User, 0),
	}
}

func (repo *UserRepository) Save(u *domain.User) error {
	if u == nil {
		return errors.New("user cannot be nil")
	}
	repo.users = append(repo.users, u)
	return nil
}

func (repo *UserRepository) FindByEmail(email string) (*domain.User, error) {
	for _, u := range repo.users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, errors.New("user not found")
}
