package repository

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type User struct {
	CreatedAt    time.Time `json:"created_at"`
	Email        string    `json:"email"`
	FirstName    string    `json:"firstname"`
	ID           uuid.UUID `json:"id"`
	LastName     string    `json:"lastname"`
	PasswordHash *string   `json:"-"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}
