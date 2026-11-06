package allowance

import (
	"errors"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamim1715/novaerp/internal/common/pagination"
	"github.com/tamim1715/novaerp/internal/common/response"
	"gorm.io/gorm"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateType(c *gin.Context) {
	var req CreateAllowanceTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	at, err := h.service.CreateType(c.Request.Context(), req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Created(c, "Allowance type created successfully", ToAllowanceTypeResponse(at))
}

func (h *Handler) FindAllTypes(c *gin.Context) {
	var req pagination.PageRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	list, total, err := h.service.FindAllTypes(c.Request.Context(), req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	req.Normalize()
	pageResp := pagination.PageResponse{
		Page:       req.Page,
		Size:       req.Size,
		TotalItems: total,
		TotalPages: int(math.Ceil(float64(total) / float64(req.Size))),
		Data:       ToAllowanceTypeResponseList(list),
	}

	response.Success(c, "Allowance types fetched successfully", pageResp)
}

func (h *Handler) FindTypeByID(c *gin.Context) {
	id := c.Param("id")
	at, err := h.service.FindTypeByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "Allowance type not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, "Allowance type fetched successfully", ToAllowanceTypeResponse(at))
}

func (h *Handler) UpdateType(c *gin.Context) {
	id := c.Param("id")
	var req UpdateAllowanceTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	at, err := h.service.UpdateType(c.Request.Context(), id, req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, "Allowance type updated successfully", ToAllowanceTypeResponse(at))
}

func (h *Handler) DeleteType(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteType(c.Request.Context(), id); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, "Allowance type deleted successfully", nil)
}

func (h *Handler) AssignAllowance(c *gin.Context) {
	var req AssignAllowanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	ea, err := h.service.AssignAllowance(c.Request.Context(), req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Created(c, "Allowance assigned to employee successfully", ToEmployeeAllowanceResponse(ea))
}

func (h *Handler) FindEmployeeAllowances(c *gin.Context) {
	var req pagination.PageRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	list, total, err := h.service.FindEmployeeAllowances(c.Request.Context(), req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	req.Normalize()
	pageResp := pagination.PageResponse{
		Page:       req.Page,
		Size:       req.Size,
		TotalItems: total,
		TotalPages: int(math.Ceil(float64(total) / float64(req.Size))),
		Data:       ToEmployeeAllowanceResponseList(list),
	}

	response.Success(c, "Employee allowance assignments fetched successfully", pageResp)
}

func (h *Handler) DeleteEmployeeAllowance(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteEmployeeAllowance(c.Request.Context(), id); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, "Employee allowance assignment deleted successfully", nil)
}
