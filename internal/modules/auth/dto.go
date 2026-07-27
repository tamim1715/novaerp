package auth

type LoginRequest struct {
	Email    string `json:"Email" validate:"required,email"`
	Password string `json:"Password" validate:"required,password"`
}
