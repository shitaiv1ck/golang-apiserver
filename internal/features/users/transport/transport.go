package users_transport

import (
	"apiserver/internal/core/domains"
	core_request "apiserver/internal/core/transport/request"
	core_response "apiserver/internal/core/transport/response"
	core_utils "apiserver/internal/core/transport/utils"
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
		responseHandler := core_response.NewResponseHandler(w)

		var userDTO UserDTO
		if err := core_request.DecodeAndValidate(r, &userDTO); err != nil {
			responseHandler.ErrorResponse("decode and validate request", err)

			return
		}

		user := domains.NewUninitializedUser(userDTO.Email, userDTO.Password)

		userDomain, err := t.usersService.Create(user)
		if err != nil {
			responseHandler.ErrorResponse("create user", err)

			return
		}

		response := CreateUserResponse{
			ID:    userDomain.ID,
			Email: userDomain.Email,
		}

		responseHandler.JsonResponse(response, http.StatusCreated)
	}
}

func (t *UsersTransport) FindByEmailHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responseHandler := core_response.NewResponseHandler(w)

		email, err := core_utils.GetStringPathValue(r, "email")
		if err != nil {
			responseHandler.ErrorResponse("get string path value", err)

			return
		}

		userDomain, err := t.usersService.FindByEmail(email)
		if err != nil {
			responseHandler.ErrorResponse("find user by email", err)

			return
		}

		response := CreateUserResponse{
			ID:    userDomain.ID,
			Email: userDomain.Email,
		}

		responseHandler.JsonResponse(response, http.StatusOK)
	}
}
