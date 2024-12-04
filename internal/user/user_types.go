package user

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID         uuid.UUID `json:"id"`
	Email      string    `json:"email"`
	Username   string    `json:"username"`
	Password   password  `json:"-"`
	Created_at time.Time `json:"created_at"`
}

type password struct {
	plaintext *string
	hash      string
}

type UserModel struct {
	DB *sql.DB
}

type UserID struct {
	Id uuid.UUID
}

var (
	ErrDuplicateEmail = errors.New("duplicate email")
)
