package irepositories

import (
	"github.com/joaofilippe/pegtech/internal/domain/entities"
)

// ClientRepository defines the interface for client operations
type ClientRepository interface {
	SaveClient(client *entities.Client) error
	GetClient(id string) (*entities.Client, error)
	GetClientByEmail(email string) (*entities.Client, error)
	ListClients() ([]*entities.Client, error)
	DeleteClient(id string) error
}
