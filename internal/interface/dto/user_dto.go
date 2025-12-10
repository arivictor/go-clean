package dto

// CreateUserRequest represents the input data for creating a user
type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UserResponse represents the output data for a user
type UserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
