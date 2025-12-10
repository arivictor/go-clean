package main

import (
	"example/internal/interface/api"
	"fmt"
)

func main() {
	// Infrastructure Layer: Initialize repository
	name := "Ari"
	email := "ari@ari.com"
	response := api.CreatUserRoute(api.CreateUserRequest{Name: name, Email: email})

	fmt.Println("Second user created successfully!", response)
}
