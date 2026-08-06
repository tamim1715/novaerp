package auth

import (
	"github.com/tamim1715/novaerp/internal/common/response"
)

var (
	_ response.APIResponse
)

// Login godoc
// @Summary User login
// @Description Authenticate user and return RS256 Access Token and Refresh Token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Login credentials"
// @Success 200 {object} response.APIResponse
// @Failure 401 {object} response.APIResponse
// @Router /auth/login [post]
func _() {}

// RefreshToken godoc
// @Summary Refresh access token
// @Description Exchange a valid refresh token for a new RS256 access token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body RefreshTokenRequest true "Refresh token payload"
// @Success 200 {object} response.APIResponse
// @Failure 401 {object} response.APIResponse
// @Router /auth/refresh [post]
func _() {}

// Logout godoc
// @Summary User logout
// @Description Revoke refresh token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body LogoutRequest true "Logout payload"
// @Success 200 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /auth/logout [post]
func _() {}
