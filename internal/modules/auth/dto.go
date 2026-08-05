package auth

import "github.com/tamim1715/novaerp/internal/modules/user"

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string            `json:"token"`
	User  user.UserResponse `json:"user"`
}
