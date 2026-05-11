package users_service

import (
	"apiserver/internal/core/domains"
	core_errors "apiserver/internal/core/errors"
	"fmt"
)

type UsersRepository interface {
	Create(user *domains.User) error
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

func (s *UsersService) Create(user *domains.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if err := user.EncryptPassword(); err != nil {
		return err
	}

	err := s.usersRepository.Create(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *UsersService) FindByEmail(email string) (*domains.User, error) {
	if len([]rune(email)) == 0 {
		return nil, fmt.Errorf("email can't be null: %w", core_errors.ErrInvalidArgument)
	}

	userFromRep, err := s.usersRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return userFromRep, nil
}
