package leaverequest

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

func (h *Handler) CreateLeaveRequest(c *gin.Context) {
	var req CreateLeaveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	lr, err := h.service.CreateLeaveRequest(c.Request.Context(), req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Created(c, "Leave request submitted successfully", ToLeaveRequestResponse(lr))
}

func (h *Handler) FindAllLeaveRequests(c *gin.Context) {
	var req pagination.PageRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	list, total, err := h.service.FindAllLeaveRequests(c.Request.Context(), req)
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
		Data:       ToLeaveRequestResponseList(list),
	}

	response.Success(c, "Leave requests fetched successfully", pageResp)
}

func (h *Handler) FindLeaveRequestByID(c *gin.Context) {
	id := c.Param("id")
	lr, err := h.service.FindLeaveRequestByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "Leave request not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, "Leave request fetched successfully", ToLeaveRequestResponse(lr))
}

func (h *Handler) UpdateLeaveStatus(c *gin.Context) {
	id := c.Param("id")
	var req UpdateLeaveStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	lr, err := h.service.UpdateLeaveStatus(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "Leave request not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, "Leave request status updated successfully", ToLeaveRequestResponse(lr))
}
