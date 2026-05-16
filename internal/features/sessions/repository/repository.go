package sessions_repository

import (
	"apiserver/internal/core/domains"
	core_errors "apiserver/internal/core/errors"
	"apiserver/internal/core/repository/postgres"
	"fmt"

	"github.com/lib/pq"
)

type SessionsRepository struct {
	store *postgres.Store
}

func NewRepository(store *postgres.Store) *SessionsRepository {
	return &SessionsRepository{
		store: store,
	}
}

func (r *SessionsRepository) Create(session *domains.Session) (*domains.Session, error) {
	db := r.store.GetDB()

	query := `
		INSERT INTO apiserver.users_sessions(session_token, csrf_token, user_id, expires_at)
		VALUES($1, $2, $3, $4)
		RETURNING session_token, csrf_token, user_id, expires_at;
	`

	sessionDomain := &domains.Session{}
	if err := db.QueryRow(
		query,
		session.SessionToken,
		session.CSRFToken,
		session.UserID,
		session.ExpiresAt,
	).Scan(
		&sessionDomain.SessionToken,
		&sessionDomain.CSRFToken,
		&sessionDomain.UserID,
		&sessionDomain.ExpiresAt,
	); err != nil {
		if errPQ, ok := err.(*pq.Error); ok {
			if errPQ.Code == "23505" {
				return nil, fmt.Errorf("%v: %w", err, core_errors.ErrConflict)
			}
		}

		return nil, err
	}

	return sessionDomain, nil
}
