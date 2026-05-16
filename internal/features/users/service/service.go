package users_service

import (
	"apiserver/internal/core/domains"
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
	if err := validateEmail(email); err != nil {
		return nil, err
	}

	userFromRep, err := s.usersRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return userFromRep, nil
}

func validateEmail(email string) error {
	if _, err := mail.ParseAddress(email); err != nil {
		return err
	}

	return nil
}
