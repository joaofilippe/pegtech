package common

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joaofilippe/pegtech/internal/domain/entities"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token has expired")
)

// ValidateToken validates a JWT token and returns the claims if valid
func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	// Check if token is valid
	if !token.Valid {
		return nil, ErrInvalidToken
	}

	// Get claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrInvalidToken
	}

	return claims, nil
}

// GetUserTypeFromClaims extracts the user type from JWT claims
func GetUserTypeFromClaims(claims jwt.MapClaims) (entities.UserType, error) {
	// Get user type from claims
	userType, ok := claims["type"].(string)
	if !ok {
		return "", ErrInvalidToken
	}

	return entities.UserType(userType), nil
}

// ValidateTokenAndGetUserType validates a JWT token and returns the user type
func ValidateTokenAndGetUserType(tokenString string) (entities.UserType, error) {
	claims, err := ValidateToken(tokenString)
	if err != nil {
		return "", err
	}

	return GetUserTypeFromClaims(claims)
}
