package api

import (
	"fmt"
	"log"

	application "example/internal/application/user"
	infrastructure "example/internal/infrastructure/persistence/inmemory"
)

// DTOs

// CreateUserRequest represents the input data for creating a user
type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UserResponse represents the output data for a user
type CreatUserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// This is just pretend, not real routing code
func CreatUserRoute(request CreateUserRequest) CreatUserResponse {
	userRepo := infrastructure.NewUserRepository()

	// Application Layer: Initialize use case with dependencies
	createUserUC := application.NewCreateUserUseCase(userRepo)

	// Execute use case
	err := createUserUC.Execute("John Doe", "john@example.com")
	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}

	fmt.Println("User created successfully!")

	return CreatUserResponse{
		Name:  request.Name,
		Email: request.Email,
	}
}
