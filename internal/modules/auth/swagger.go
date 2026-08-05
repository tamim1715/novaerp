package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/tamim1715/novaerp/internal/common/response"
)

var (
	_ response.APIResponse
)

// LoginDoc godoc
// @Summary User Login
// @Description Authenticate user and return JWT bearer token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Login Credentials"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 401 {object} response.APIResponse
// @Router /auth/login [post]
func (h *Handler) LoginDoc(c *gin.Context) {
	h.Login(c)
}
