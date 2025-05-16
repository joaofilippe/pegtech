package userusecases

import (
	"time"

	"github.com/joaofilippe/pegtech/domain/entities"
	"github.com/joaofilippe/pegtech/domain/irepositories"
)

// UpdateUserInput defines the input for updating a user
type UpdateUserInput struct {
	ID       string
	Username string
	Email    string
}

// UpdateUserCase handles user information updates
type UpdateUserCase struct {
	userRepo irepositories.UserRepository
}

// NewUpdateUserCase creates a new instance of UpdateUserCase
func NewUpdateUserCase(userRepo irepositories.UserRepository) *UpdateUserCase {
	return &UpdateUserCase{
		userRepo: userRepo,
	}
}

// Execute performs the user update operation
func (uc *UpdateUserCase) Execute(input UpdateUserInput) (*entities.User, error) {
	// Validate input
	if input.Username == "" || input.Email == "" {
		return nil, ErrInvalidUserData
	}

	// Get existing user
	user, err := uc.userRepo.GetUserByID(input.ID)
	if err != nil {
		return nil, ErrUserNotFound
	}

	// Update user information
	user.Username = input.Username
	user.Email = input.Email
	user.UpdatedAt = time.Now()

	// Save updated user
	if err := uc.userRepo.SaveUser(user); err != nil {
		return nil, err
	}

	return user, nil
}
