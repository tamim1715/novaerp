package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamim1715/novaerp/internal/common/response"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.service.Login(c.Request.Context(), req)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	response.Success(c, "Login successful", resp)
}
