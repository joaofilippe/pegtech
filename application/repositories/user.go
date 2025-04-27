package irepositories

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/pegtech/application/entities"
)

type IUserRepository interface {
	Save(email string, password string) error
	FindByID(uuid uuid.UUID) entities.User
	FindByEmail(email string) entities.User
	Update(email string, password string) error
	Delete(uuid uuid.UUID) error
}
