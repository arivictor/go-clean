package inmemory

import (
	"errors"
	"example/internal/domain/user"
)

// UserRepository is an in-memory implementation of user.Repository
type UserRepository struct {
	users []*user.User
}

// NewUserRepository creates a new in-memory user repository
func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make([]*user.User, 0),
	}
}

// Save stores a user in memory
func (repo *UserRepository) Save(u *user.User) error {
	if u == nil {
		return errors.New("user cannot be nil")
	}
	repo.users = append(repo.users, u)
	return nil
}

// FindByEmail finds a user by email address
func (repo *UserRepository) FindByEmail(email string) (*user.User, error) {
	for _, u := range repo.users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, errors.New("user not found")
}
