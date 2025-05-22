package userusecases

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joaofilippe/pegtech/internal/domain/entities"
	"github.com/joaofilippe/pegtech/internal/domain/irepositories"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

// LoginUserInput defines the input for user login
type LoginUserInput struct {
	Email    string
	Password string
}

// LoginResponse defines the response for successful login
type LoginResponse struct {
	Token     string
	UserID    string
	UserType  entities.UserType
	ExpiresAt time.Time
}

// LoginUserCase handles user authentication
type LoginUserCase struct {
	userRepo irepositories.UserRepository
	jwtKey   []byte
}

// NewLoginUserCase creates a new instance of LoginUserCase
func NewLoginUserCase(userRepo irepositories.UserRepository) *LoginUserCase {
	return &LoginUserCase{
		userRepo: userRepo,
		jwtKey:   []byte(os.Getenv("JWT_SECRET")),
	}
}

// Execute performs the login operation
func (uc *LoginUserCase) Execute(input LoginUserInput) (*LoginResponse, error) {
	// Get user by email
	user, err := uc.userRepo.GetUserByEmail(input.Email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	// Calculate expiration time (72 hours from now)
	expiresAt := time.Now().Add(72 * time.Hour)

	// Create JWT token
	claims := jwt.MapClaims{
		"user_id": user.ID.String(),
		"type":    user.Type,
		"exp":     expiresAt.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token
	tokenString, err := token.SignedString(uc.jwtKey)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		Token:     tokenString,
		UserID:    user.ID.String(),
		UserType:  user.Type,
		ExpiresAt: expiresAt,
	}, nil
}
