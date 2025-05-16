package repositories

import (
	"github.com/joaofilippe/pegtech/internal/domain/entities"
	"github.com/joaofilippe/pegtech/internal/domain/irepositories"
	"github.com/joaofilippe/pegtech/internal/infra/repositories/database"
)

// UserRepository implements the UserRepository interface using PostgreSQL
type UserRepository struct {
	db *database.PostgresDB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *database.PostgresDB) irepositories.UserRepository {
	return &UserRepository{
		db: db,
	}
}

// SaveUser saves a user to the database
func (r *UserRepository) SaveUser(user *entities.User) error {
	db := r.db.DB()

	query := `
		INSERT INTO users (id, username, email, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (id) DO UPDATE
		SET username = $2, email = $3, password = $4, updated_at = $6
	`

	_, err := db.Exec(query,
		user.ID,
		user.Username,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
	)

	return err
}

// GetUserByEmail retrieves a user by their email
func (r *UserRepository) GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	db := r.db.DB()

	query := `SELECT * FROM users WHERE email = $1`

	err := db.Get(&user, query, email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUserByID retrieves a user by their ID
func (r *UserRepository) GetUserByID(id string) (*entities.User, error) {
	var user entities.User
	db := r.db.DB()
	query := `SELECT * FROM users WHERE id = $1`

	err := db.Get(&user, query, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// DeleteUser removes a user from the database
func (r *UserRepository) DeleteUser(id string) error {
	db := r.db.DB()
	query := `DELETE FROM users WHERE id = $1`

	_, err := db.Exec(query, id)
	return err
}
