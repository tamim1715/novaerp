package user

import (
	"errors"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamim1715/novaerp/internal/common/pagination"
	"github.com/tamim1715/novaerp/internal/common/response"
	"github.com/tamim1715/novaerp/internal/common/validator"
	"gorm.io/gorm"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Create(c *gin.Context) {

	var request CreateUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := validator.Validate.Struct(request); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.service.Create(c.Request.Context(), request)

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Created(
		c,
		"User created successfully",
		ToResponse(user),
	)
}

func (h *Handler) FindAll(c *gin.Context) {

	var request pagination.PageRequest

	if err := c.ShouldBindQuery(&request); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	users, total, err := h.service.FindAll(c.Request.Context(), request)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	request.Normalize()

	pageResponse := pagination.PageResponse{
		Page:       request.Page,
		Size:       request.Size,
		TotalItems: total,
		TotalPages: int(math.Ceil(float64(total) / float64(request.Size))),
		Data:       ToResponses(users),
	}

	response.Success(
		c,
		"Users fetched successfully",
		pageResponse,
	)
}

func (h *Handler) FindByID(c *gin.Context) {

	id := c.Param("id")

	user, err := h.service.FindByID(c.Request.Context(), id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "User not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(
		c,
		"User fetched successfully",
		ToResponse(user),
	)
}

func (h *Handler) Update(c *gin.Context) {

	id := c.Param("id")

	var request UpdateUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := validator.Validate.Struct(request); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.service.Update(c.Request.Context(), id, request)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "User not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(
		c,
		"User updated successfully",
		ToResponse(user),
	)
}

func (h *Handler) Delete(c *gin.Context) {

	id := c.Param("id")

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "User not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(
		c,
		"User deleted successfully",
		nil,
	)
}
