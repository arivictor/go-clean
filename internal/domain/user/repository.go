package user

// Repository defines the interface for User persistence
// This belongs to the domain layer but implementations are in infrastructure
type Repository interface {
	Save(user *User) error
	FindByEmail(email string) (*User, error)
}
