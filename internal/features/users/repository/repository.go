package users_repository

import (
	"apiserver/internal/core/domains"
	core_errors "apiserver/internal/core/errors"
	"apiserver/internal/core/repository/postgres"
	"database/sql"
	"errors"
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
		RETURNING id
	`

	if err := db.QueryRow(query, user.Email, user.EncryptedPassword).Scan(&user.ID); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UsersRepository) FindByEmail(email string) (*domains.User, error) {
	db := r.Store.GetDB()

	query := `
		SELECT * FROM apiserver.users
		WHERE email = $1
	`

	user := &domains.User{}
	if err := db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.EncryptedPassword); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, core_errors.ErrNotFound
		}

		return nil, err
	}

	return user, nil
}
