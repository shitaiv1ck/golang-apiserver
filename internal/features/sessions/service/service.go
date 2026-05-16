package sessions_service

import (
	"apiserver/internal/core/domains"
	core_errors "apiserver/internal/core/errors"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"net/mail"
	"time"
)

type UsersRepository interface {
	FindByEmail(email string) (*domains.User, error)
}

type SessionsRepository interface {
	Create(session *domains.Session) (*domains.Session, error)
}

type SessionService struct {
	usersRepository    UsersRepository
	sessionsRepository SessionsRepository
}

func NewService(usersRepository UsersRepository, sessionsRepository SessionsRepository) *SessionService {
	return &SessionService{
		usersRepository:    usersRepository,
		sessionsRepository: sessionsRepository,
	}
}

func (s *SessionService) Authenticate(email string, password string) (int, error) {
	if err := validateEmail(email); err != nil {
		return -1, core_errors.ErrInvalidPasswordOrEmail
	}

	if err := validatePassword(password); err != nil {
		return -1, core_errors.ErrInvalidPasswordOrEmail
	}

	user, err := s.usersRepository.FindByEmail(email)
	if err != nil {
		if errors.Is(err, core_errors.ErrNotFound) {
			return -1, core_errors.ErrInvalidPasswordOrEmail
		}

		return -1, err
	}

	if err := user.VerifyPassword(password); err != nil {
		return -1, core_errors.ErrInvalidPasswordOrEmail
	}

	return user.ID, nil
}

func (s *SessionService) Create(userID int) (*domains.Session, error) {
	sessionToken, err := generateToken(32)
	if err != nil {
		return nil, err
	}

	csrfToken, err := generateToken(32)
	if err != nil {
		return nil, err
	}

	session := domains.NewSession(sessionToken, csrfToken, userID, time.Now().Add(24*time.Hour))

	sessionFromRep, err := s.sessionsRepository.Create(session)
	if err != nil {
		return nil, err
	}

	return sessionFromRep, nil
}

func validateEmail(email string) error {
	if _, err := mail.ParseAddress(email); err != nil {
		return err
	}

	return nil
}

func validatePassword(password string) error {
	if len([]rune(password)) == 0 {
		return core_errors.ErrInvalidArgument
	}

	return nil
}

func generateToken(len int) (string, error) {
	bytes := make([]byte, len)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	token := base64.URLEncoding.EncodeToString(bytes)

	return token, nil
}
