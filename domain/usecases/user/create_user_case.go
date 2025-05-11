package userusecases

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/joaofilippe/pegtech/domain/entities"
	"github.com/joaofilippe/pegtech/domain/irepositories"
)

// CreateUserInput defines the input for creating a new user
type CreateUserInput struct {
	Username string
	Email    string
	Password string
}

// CreateUserCase handles user creation
type CreateUserCase struct {
	userRepo irepositories.UserRepository
}

// NewCreateUserCase creates a new instance of CreateUserCase
func NewCreateUserCase(userRepo irepositories.UserRepository) *CreateUserCase {
	return &CreateUserCase{
		userRepo: userRepo,
	}
}

// Execute performs the user creation operation
func (uc *CreateUserCase) Execute(input CreateUserInput) (*entities.User, error) {
	// Validate input
	if input.Username == "" || input.Email == "" || input.Password == "" {
		return nil, ErrInvalidUserData
	}

	// Check if user already exists
	existingUser, err := uc.userRepo.GetUserByEmail(input.Email)
	if err == nil && existingUser != nil {
		return nil, ErrUserAlreadyExists
	}

	// Create new user
	user := &entities.User{
		ID:        generateID(),
		Username:  input.Username,
		Email:     input.Email,
		Password:  input.Password, // Note: In a real application, this should be hashed
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Save user
	if err := uc.userRepo.SaveUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

// Helper function to generate unique IDs
func generateID() string {
	bytes := make([]byte, 8)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
