package users_repository

import (
	"apiserver/internal/core/domains"
	core_errors "apiserver/internal/core/errors"
	"apiserver/internal/core/repository/postgres"
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
)

type UsersRepository struct {
	Store *postgres.Store
}

func NewRepository(store *postgres.Store) *UsersRepository {
	return &UsersRepository{
		Store: store,
	}
}

func (r *UsersRepository) Create(user *domains.User) (*domains.User, error) {
	db := r.Store.GetDB()

	query := `
		INSERT INTO apiserver.users(email, encrypted_password)
		VALUES ($1, $2)
		RETURNING id, email;
	`

	userDomain := &domains.User{}
	if err := db.QueryRow(query, user.Email, user.EncryptedPassword).Scan(&userDomain.ID, &userDomain.Email); err != nil {
		if errPQ, ok := err.(*pq.Error); ok {
			if errPQ.Code == "23505" {
				return nil, fmt.Errorf("%v: %w", err, core_errors.ErrInvalidArgument)
			}
		}

		return nil, err
	}

	return userDomain, nil
}

func (r *UsersRepository) FindByEmail(email string) (*domains.User, error) {
	db := r.Store.GetDB()

	query := `
		SELECT id, email FROM apiserver.users
		WHERE email = $1;
	`

	user := &domains.User{}
	if err := db.QueryRow(query, email).Scan(&user.ID, &user.Email); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, core_errors.ErrNotFound
		}

		return nil, err
	}

	return user, nil
}
