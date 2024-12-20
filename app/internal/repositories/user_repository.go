package repositories

import (
	"app/internal/models"
	"app/pkg/errors"
	"database/sql"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUser(email string) (*models.User, error)
	IsEmailExist(email string) (bool, error)
}

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) CreateUser(user *models.User) error {
	query := `
		INSERT INTO users(name, email, password, created_at) VALUES (?, ?, ?, ?)
	`

	result, err := r.DB.Exec(
		query,
		user.Name,
		user.Email,
		user.Password,
		user.CreatedAt,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = id

	return nil
}

func (r *userRepository) IsEmailExist(email string) (bool, error) {
	query := "SELECT id FROM users WHERE email = ? LIMIT 1"

	var id int64

	err := r.DB.QueryRow(query, email).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (r *userRepository) GetUser(email string) (*models.User, error) {
	user := &models.User{}

	query := "SELECT id, name, email, password FROM users WHERE email = ? LIMIT 1"

	err := r.DB.QueryRow(
		query,
		email,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrUserNotExist
		}

		return nil, err
	}

	return user, nil
}
