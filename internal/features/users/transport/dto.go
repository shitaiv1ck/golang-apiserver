package users_transport

type UserDTO struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CreateUserResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}
