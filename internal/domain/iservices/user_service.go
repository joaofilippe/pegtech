package iservices

import (
	"github.com/joaofilippe/pegtech/internal/domain/entities"
	userusecases "github.com/joaofilippe/pegtech/internal/domain/usecases/user"
)

// UserService defines the interface for user operations
type UserService interface {
	// CreateUser creates a new user in the system
	CreateUser(username, name, email, password, phone string, userType entities.UserType) (*entities.User, error)

	// GetUserByEmail retrieves a user by their email
	GetUserByEmail(email string) (*entities.User, error)

	// GetUserByID retrieves a user by their ID
	GetUserByID(id string) (*entities.User, error)

	// UpdateUser updates an existing user's information
	UpdateUser(id string, username, email string) (*entities.User, error)

	// DeleteUser removes a user from the system
	DeleteUser(id string) error

	// Login authenticates a user and returns a JWT token
	Login(email, password string) (*userusecases.LoginResponse, error)
}
