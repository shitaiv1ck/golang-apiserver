package users_transport

import (
	"apiserver/internal/core/domains"
	core_errors "apiserver/internal/core/errors"
	core_request "apiserver/internal/core/request"
	core_utils "apiserver/internal/core/utils"
	"encoding/json"
	"errors"
	"net/http"
)

type UsersService interface {
	Create(user *domains.User) (*domains.User, error)
	FindByEmail(email string) (*domains.User, error)
}

type UsersTransport struct {
	usersService UsersService
}

func NewTransport(usersService UsersService) *UsersTransport {
	return &UsersTransport{
		usersService: usersService,
	}
}

func (t *UsersTransport) CreateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userDTO UserDTO
		if err := core_request.DecodeAndValidate(r, &userDTO); err != nil {
			if errors.Is(err, core_errors.ErrInvalidArgument) {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}

			errorDTO := ErrorDTO{
				Message: err.Error(),
			}

			json.NewEncoder(w).Encode(errorDTO)

			return
		}

		userFromDTO := domains.NewUser(0, userDTO.Email, userDTO.Password, "")

		user, err := t.usersService.Create(userFromDTO)
		if err != nil {
			if errors.Is(err, core_errors.ErrInvalidArgument) {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}

			errorDTO := ErrorDTO{
				Message: err.Error(),
			}

			json.NewEncoder(w).Encode(errorDTO)

			return
		}

		response := CreateUserResponse{
			ID:                user.ID,
			Email:             user.Email,
			EncryptedPassword: user.EncryptedPassword,
		}

		json.NewEncoder(w).Encode(response)
	}
}

func (t *UsersTransport) FindByEmailHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email, err := core_utils.GetStringPathValue(r, "email")
		if err != nil {
			errorDTO := ErrorDTO{
				Message: err.Error(),
			}

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errorDTO)

			return
		}

		userDomain, err := t.usersService.FindByEmail(email)
		if err != nil {
			if errors.Is(err, core_errors.ErrNotFound) {
				w.WriteHeader(http.StatusNotFound)
			} else if errors.Is(err, core_errors.ErrInvalidArgument) {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}

			errorDTO := ErrorDTO{
				Message: err.Error(),
			}

			json.NewEncoder(w).Encode(errorDTO)
			return
		}

		response := CreateUserResponse{
			ID:                userDomain.ID,
			Email:             userDomain.Email,
			EncryptedPassword: userDomain.EncryptedPassword,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
