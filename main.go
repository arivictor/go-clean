package main

import (
	"fmt"
	"log"

	"example/internal/application/usecase"
	"example/internal/infrastructure/persistence/inmemory"
)

func main() {
	// Infrastructure Layer: Initialize repository
	userRepo := inmemory.NewUserRepository()

	// Application Layer: Initialize use case with dependencies
	createUserUC := usecase.NewCreateUserUseCase(userRepo)

	// Execute use case
	err := createUserUC.Execute("John Doe", "john@example.com")
	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}

	fmt.Println("User created successfully!")

	// Try to create an invalid user
	err = createUserUC.Execute("", "invalid")
	if err != nil {
		fmt.Printf("Validation error (expected): %v\n", err)
	}

	// Create another valid user
	err = createUserUC.Execute("Jane Smith", "jane@example.com")
	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}

	fmt.Println("Second user created successfully!")
}
