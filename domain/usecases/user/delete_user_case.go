package userusecases

import (
	"github.com/joaofilippe/pegtech/domain/irepositories"
)

// DeleteUserCase handles user deletion
type DeleteUserCase struct {
	userRepo irepositories.UserRepository
}

// NewDeleteUserCase creates a new instance of DeleteUserCase
func NewDeleteUserCase(userRepo irepositories.UserRepository) *DeleteUserCase {
	return &DeleteUserCase{
		userRepo: userRepo,
	}
}

// Execute performs the user deletion operation
func (uc *DeleteUserCase) Execute(id string) error {
	// Check if user exists
	_, err := uc.userRepo.GetUserByID(id)
	if err != nil {
		return ErrUserNotFound
	}

	return uc.userRepo.DeleteUser(id)
}
