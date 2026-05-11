package domains

import (
	core_errors "apiserver/internal/core/errors"
	"fmt"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                int
	Email             string
	Password          string
	EncryptedPassword string
}

func NewUser(id int, email string, password string, encryptedPassword string) *User {
	return &User{
		ID:                id,
		Email:             email,
		Password:          password,
		EncryptedPassword: encryptedPassword,
	}
}

func (u *User) Validate() error {
	if len([]rune(u.Email)) == 0 {
		return fmt.Errorf("email can't be null: %w", core_errors.ErrInvalidArgument)
	}

	if _, err := mail.ParseAddress(u.Email); err != nil {
		return fmt.Errorf("invalid 'email' format: %v: %w", err, core_errors.ErrInvalidArgument)
	}

	if len([]rune(u.Password)) == 0 {
		return fmt.Errorf("password can't be null: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}

func (u *User) EncryptPassword() error {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		return fmt.Errorf("generate encrypt password: %w", err)
	}

	u.EncryptedPassword = string(encryptedPassword)

	return nil
}
