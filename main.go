package main

import (
	"example/internal/interface/api"
	"fmt"
)

func main() {
	name := "Ari"
	email := "ari@ari.com"
	response := api.CreatUserRoute(api.CreateUserRequest{Name: name, Email: email})

	fmt.Println("Second user created successfully!", response)
}
