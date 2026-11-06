package shift

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

func (h *Handler) CreateShift(c *gin.Context) {
	var req CreateShiftRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	sh, err := h.service.CreateShift(c.Request.Context(), req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Created(c, "Shift created successfully", ToShiftResponse(sh))
}

func (h *Handler) FindAllShifts(c *gin.Context) {
	var req pagination.PageRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	list, total, err := h.service.FindAllShifts(c.Request.Context(), req)
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
		Data:       ToShiftResponseList(list),
	}

	response.Success(c, "Shifts fetched successfully", pageResp)
}

func (h *Handler) FindShiftByID(c *gin.Context) {
	id := c.Param("id")
	sh, err := h.service.FindShiftByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "Shift not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, "Shift fetched successfully", ToShiftResponse(sh))
}

func (h *Handler) UpdateShift(c *gin.Context) {
	id := c.Param("id")
	var req UpdateShiftRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	sh, err := h.service.UpdateShift(c.Request.Context(), id, req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, "Shift updated successfully", ToShiftResponse(sh))
}

func (h *Handler) DeleteShift(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteShift(c.Request.Context(), id); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, "Shift deleted successfully", nil)
}

func (h *Handler) AssignShift(c *gin.Context) {
	var req AssignShiftRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	es, err := h.service.AssignShift(c.Request.Context(), req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Created(c, "Shift assigned to employee successfully", ToEmployeeShiftResponse(es))
}

func (h *Handler) FindEmployeeShifts(c *gin.Context) {
	var req pagination.PageRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	list, total, err := h.service.FindEmployeeShifts(c.Request.Context(), req)
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
		Data:       ToEmployeeShiftResponseList(list),
	}

	response.Success(c, "Employee shift assignments fetched successfully", pageResp)
}
