package department

import (
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamim1715/novaerp/internal/common/pagination"
	"github.com/tamim1715/novaerp/internal/common/response"
	"github.com/tamim1715/novaerp/internal/common/validator"
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

	var request CreateDepartmentRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := validator.Validate.Struct(request); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	department, err := h.service.Create(request)

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Created(
		c,
		"Department created successfully",
		ToResponse(department),
	)
}

func (h *Handler) FindAll(c *gin.Context) {

	var request pagination.PageRequest

	if err := c.ShouldBindQuery(&request); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	departments, total, err := h.service.FindAll(request)

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
		Data:       ToResponses(departments),
	}

	response.Success(
		c,
		"Departments fetched successfully",
		pageResponse,
	)
}

func (h *Handler) FindByID(c *gin.Context) {

	id := c.Param("id")

	department, err := h.service.FindByID(id)

	if err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return
	}

	response.Success(
		c,
		"Department fetched successfully",
		ToResponse(department),
	)
}

func (h *Handler) Update(c *gin.Context) {

	id := c.Param("id")

	var request UpdateDepartmentRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := validator.Validate.Struct(request); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	department, err := h.service.Update(id, request)

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(
		c,
		"Department updated successfully",
		ToResponse(department),
	)
}

func (h *Handler) Delete(c *gin.Context) {

	id := c.Param("id")

	if err := h.service.Delete(id); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(
		c,
		"Department deleted successfully",
		nil,
	)
}
