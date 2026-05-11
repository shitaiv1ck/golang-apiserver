package users_transport

type UserDTO struct {
	Email             string `json:"email" validate:"required"`
	Password          string `json:"password" validate:"required"`
}

type ErrorDTO struct {
	Message string `json:"message"`
}

type CreateUserResponse struct {
	ID                int    `json:"id"`
	Email             string `json:"email"`
	EncryptedPassword string `json:"encrypted_password"`
}
