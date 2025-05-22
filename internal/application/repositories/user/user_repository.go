package repositories

import (
	"database/sql"
	"errors"

	"github.com/joaofilippe/pegtech/internal/domain/entities"
	"github.com/joaofilippe/pegtech/internal/domain/irepositories"
	"github.com/joaofilippe/pegtech/internal/infra/repositories/database"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

// UserRepository implements the UserRepository interface
type UserRepository struct {
	db *database.PostgresDB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *database.PostgresDB) irepositories.UserRepository {
	return &UserRepository{
		db: db,
	}
}

// SaveUser saves a user to the storage
func (r *UserRepository) SaveUser(user *entities.User) error {
	_, err := r.db.DB().Exec(SaveUserQuery,
		user.ID,
		user.Name,
		user.Email,
		user.Username,
		user.Phone,
		user.Password,
		user.Type,
		user.CreatedAt,
		user.UpdatedAt,
	)

	return err
}

// GetUser retrieves a user by ID
func (r *UserRepository) GetUser(id string) (*entities.User, error) {
	user := &entities.User{}
	err := r.db.DB().QueryRow(GetUserQuery, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Type,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, ErrUserNotFound
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByEmail retrieves a user by email
func (r *UserRepository) GetUserByEmail(email string) (*entities.User, error) {
	user := &entities.User{}
	err := r.db.DB().QueryRow(GetUserByEmailQuery, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Type,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, ErrUserNotFound
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

// ListUsers retrieves all users
func (r *UserRepository) ListUsers() ([]*entities.User, error) {
	rows, err := r.db.DB().Query(ListUsersQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entities.User
	for rows.Next() {
		user := &entities.User{}
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.Type,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// DeleteUser deletes a user by ID
func (r *UserRepository) DeleteUser(id string) error {
	result, err := r.db.DB().Exec(DeleteUserQuery, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrUserNotFound
	}

	return nil
}

// GetUserByID retrieves a user by ID
func (r *UserRepository) GetUserByID(id string) (*entities.User, error) {
	user := &entities.User{}
	err := r.db.DB().QueryRow(GetUserQuery, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Type,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, ErrUserNotFound
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}
