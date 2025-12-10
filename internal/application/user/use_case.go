package application

import (
	domain "example/internal/domain/user"
)

// CreateUserUseCase handles the creation of new users
type CreateUserUseCase struct {
	userRepo domain.Repository
}

// NewCreateUserUseCase creates a new instance of CreateUserUseCase
func NewCreateUserUseCase(userRepo domain.Repository) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepo: userRepo,
	}
}

// Execute performs the user creation use case
func (uc *CreateUserUseCase) Execute(name, email string) error {
	newUser, err := domain.NewUser(name, email)
	if err != nil {
		return err
	}

	return uc.userRepo.Save(newUser)
}
