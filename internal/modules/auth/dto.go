package auth

import "github.com/tamim1715/novaerp/internal/modules/user"

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"admin@novaerp.com"`
	Password string `json:"password" binding:"required,min=6" example:"secret123"`
}

type LoginResponse struct {
	AccessToken  string            `json:"accessToken"`
	RefreshToken string            `json:"refreshToken"`
	TokenType    string            `json:"tokenType" example:"Bearer"`
	ExpiresIn    int64             `json:"expiresIn" example:"900"` // 15 minutes in seconds
	User         user.UserResponse `json:"user"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required" example:"6f4a8b..."`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	TokenType    string `json:"tokenType" example:"Bearer"`
	ExpiresIn    int64  `json:"expiresIn" example:"900"`
}

type LogoutRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required" example:"6f4a8b..."`
}
