package domain

import (
	"errors"
	"strings"
)

// User represents the core business entity
type User struct {
	Name  string
	Email string
}

// NewUser creates a new User with validation
func NewUser(name, email string) (*User, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("name cannot be empty")
	}
	if strings.TrimSpace(email) == "" || !strings.Contains(email, "@") {
		return nil, errors.New("invalid email")
	}

	return &User{
		Name:  name,
		Email: email,
	}, nil
}
