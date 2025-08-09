package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type User struct {
	CreatedAt    time.Time `json:"created_at"`
	Email        string    `json:"email"`
	ID           uuid.UUID `json:"id"`
	PasswordHash *string   `json:"-"`
	UpdatedAt    time.Time `json:"updated_at"`
	Username     string    `json:"username"`
}
type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (userRepo *UserRepository) CreateUser(user *User) (*User, error) {
	query := `
		INSERT INTO users (username, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at
	`

	err := userRepo.db.QueryRow(query, user.Username, user.Email, user.PasswordHash).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok && pgErr.Code == "23505" {
			return nil, errors.New("email already exists")
		}
		return nil, fmt.Errorf("insert user: %w", err)
	}

	return user, nil
}

func (userRepo *UserRepository) GetUserByEmail(user *User) (*User, error) {
	query := `
	  SELECT id, username, password_hash FROM users
	  WHERE email = $1
  `

	err := userRepo.db.QueryRow(query, user.Email).Scan(&user.ID, &user.Username, &user.PasswordHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// user not found
			return nil, nil
		}
		return nil, fmt.Errorf("query login user: %w", err)
	}

	return user, nil
}
