package account

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID                uuid.UUID `json:"id"`
	Name              string    `json:"name"`
	Email             string    `json:"email"`
	Phone             string    `json:"phone"`
	Password          string    `json:"password"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	IsEmailVerified   bool      `json:"is_email_verified"`
}

type CreateAccountRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type AccountResponse struct {
	ID                uuid.UUID `json:"id"`
	Name              string    `json:"name"`
	Email             string    `json:"email"`
	Phone             string    `json:"phone"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	IsEmailVerified   bool      `json:"is_email_verified"`
}
