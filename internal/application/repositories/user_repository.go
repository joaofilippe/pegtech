package repositories

import (
	"database/sql"
	"errors"

	"github.com/joaofilippe/pegtech/internal/domain/entities"
	"github.com/joaofilippe/pegtech/internal/infra/repositories/database"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserRepository struct {
	db *database.PostgresDB
}

func NewUserRepository(db *database.PostgresDB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) SaveUser(user *entities.User) error {
	query := `
		INSERT INTO users (id, username, name, email, password, type, active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		ON CONFLICT (id) DO UPDATE
		SET username = $2, name = $3, email = $4, password = $5, type = $6, active = $7, updated_at = $9
	`

	_, err := r.db.DB().Exec(query,
		user.ID,
		user.Username,
		user.Name,
		user.Email,
		user.Password,
		user.Type,
		user.Active,
		user.CreatedAt,
		user.UpdatedAt,
	)

	return err
}

func (r *UserRepository) GetUser(id string) (*entities.User, error) {
	query := `
		SELECT id, username, name, email, password, type, active, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	user := &entities.User{}
	err := r.db.DB().QueryRow(query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Type,
		&user.Active,
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

func (r *UserRepository) GetUserByEmail(email string) (*entities.User, error) {
	query := `
		SELECT id, username, name, email, password, type, active, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	user := &entities.User{}
	err := r.db.DB().QueryRow(query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Type,
		&user.Active,
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

func (r *UserRepository) ListUsers() ([]*entities.User, error) {
	query := `
		SELECT id, username, name, email, password, type, active, created_at, updated_at
		FROM users
		ORDER BY created_at DESC
	`

	rows, err := r.db.DB().Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entities.User
	for rows.Next() {
		user := &entities.User{}
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.Type,
			&user.Active,
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

func (r *UserRepository) DeleteUser(id string) error {
	query := `
		DELETE FROM users
		WHERE id = $1
	`

	result, err := r.db.DB().Exec(query, id)
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

func (r *UserRepository) GetUserByID(id string) (*entities.User, error) {
	return r.GetUser(id)
}
