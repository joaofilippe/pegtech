package irepositories

import (
	"github.com/joaofilippe/pegtech/internal/domain/entities"
)

// UserRepository defines the interface for user operations
type UserRepository interface {
	SaveUser(user *entities.User) error
	GetUserByEmail(email string) (*entities.User, error)
	GetUserByID(id string) (*entities.User, error)
	DeleteUser(id string) error
}
