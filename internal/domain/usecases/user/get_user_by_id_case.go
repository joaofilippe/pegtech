package userusecases

import (
	"github.com/joaofilippe/pegtech/internal/domain/entities"
	"github.com/joaofilippe/pegtech/internal/domain/irepositories"
)

// GetUserByIDCase handles user retrieval by ID
type GetUserByIDCase struct {
	userRepo irepositories.UserRepository
}

// NewGetUserByIDCase creates a new instance of GetUserByIDCase
func NewGetUserByIDCase(userRepo irepositories.UserRepository) *GetUserByIDCase {
	return &GetUserByIDCase{
		userRepo: userRepo,
	}
}

// Execute performs the user retrieval by ID operation
func (uc *GetUserByIDCase) Execute(id string) (*entities.User, error) {
	return uc.userRepo.GetUserByID(id)
}
