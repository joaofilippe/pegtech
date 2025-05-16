package userusecases

import (
	"github.com/joaofilippe/pegtech/internal/domain/entities"
	"github.com/joaofilippe/pegtech/internal/domain/irepositories"
)

// GetUserByEmailCase handles user retrieval by email
type GetUserByEmailCase struct {
	userRepo irepositories.UserRepository
}

// NewGetUserByEmailCase creates a new instance of GetUserByEmailCase
func NewGetUserByEmailCase(userRepo irepositories.UserRepository) *GetUserByEmailCase {
	return &GetUserByEmailCase{
		userRepo: userRepo,
	}
}

// Execute performs the user retrieval by email operation
func (uc *GetUserByEmailCase) Execute(email string) (*entities.User, error) {
	return uc.userRepo.GetUserByEmail(email)
}
