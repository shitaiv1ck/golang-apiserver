package users_service

import (
	"apiserver/internal/core/domains"
	core_errors "apiserver/internal/core/errors"
	"fmt"
	"net/mail"
)

type UsersRepository interface {
	Create(user *domains.User) (*domains.User, error)
	FindByEmail(email string) (*domains.User, error)
}

type UsersService struct {
	usersRepository UsersRepository
}

func NewService(usersRepository UsersRepository) *UsersService {
	return &UsersService{
		usersRepository: usersRepository,
	}
}

func (s *UsersService) Create(user *domains.User) (*domains.User, error) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.EncryptPassword(); err != nil {
		return nil, err
	}

	userDomain, err := s.usersRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return userDomain, nil
}

func (s *UsersService) FindByEmail(email string) (*domains.User, error) {
	if _, err := mail.ParseAddress(email); err != nil {
		return nil, fmt.Errorf("invalid 'email' format: %v: %w", err, core_errors.ErrInvalidArgument)
	}

	userFromRep, err := s.usersRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return userFromRep, nil
}
