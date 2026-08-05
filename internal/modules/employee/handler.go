package employee

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

	var request CreateEmployeeRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := validator.Validate.Struct(request); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	employee, err := h.service.Create(c.Request.Context(), request)

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Created(
		c,
		"Employee created successfully",
		ToResponse(employee),
	)
}

func (h *Handler) FindAll(c *gin.Context) {

	var request pagination.PageRequest

	if err := c.ShouldBindQuery(&request); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	employees, total, err := h.service.FindAll(c.Request.Context(), request)

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
		Data:       ToResponses(employees),
	}

	response.Success(
		c,
		"Employees fetched successfully",
		pageResponse,
	)
}

func (h *Handler) FindByID(c *gin.Context) {

	id := c.Param("id")

	employee, err := h.service.FindByID(c.Request.Context(), id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "Employee not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(
		c,
		"Employee fetched successfully",
		ToResponse(employee),
	)
}

func (h *Handler) Update(c *gin.Context) {

	id := c.Param("id")

	var request UpdateEmployeeRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := validator.Validate.Struct(request); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	employee, err := h.service.Update(c.Request.Context(), id, request)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "Employee not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(
		c,
		"Employee updated successfully",
		ToResponse(employee),
	)
}

func (h *Handler) Delete(c *gin.Context) {

	id := c.Param("id")

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "Employee not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(
		c,
		"Employee deleted successfully",
		nil,
	)
}
