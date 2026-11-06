package attendance

import (
	"errors"
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func (h *Handler) CheckIn(c *gin.Context) {
	var req CheckInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	att, err := h.service.CheckIn(c.Request.Context(), req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, "Checked in successfully", ToAttendanceResponse(att))
}

func (h *Handler) CheckOut(c *gin.Context) {
	var req CheckOutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	att, err := h.service.CheckOut(c.Request.Context(), req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, "Checked out successfully", ToAttendanceResponse(att))
}

func (h *Handler) Create(c *gin.Context) {
	var req CreateAttendanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	att, err := h.service.CreateAttendance(c.Request.Context(), req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Created(c, "Attendance record created successfully", ToAttendanceResponse(att))
}

func (h *Handler) FindAll(c *gin.Context) {
	var req pagination.PageRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	list, total, err := h.service.FindAll(c.Request.Context(), req)
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
		Data:       ToAttendanceResponseList(list),
	}

	response.Success(c, "Attendance records fetched successfully", pageResp)
}

func (h *Handler) FindByID(c *gin.Context) {
	id := c.Param("id")
	att, err := h.service.FindByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "Attendance record not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, "Attendance record fetched successfully", ToAttendanceResponse(att))
}

func (h *Handler) ProcessDaily(c *gin.Context) {
	var req ProcessDailyAttendanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.service.ProcessDailyAttendance(c.Request.Context(), req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, "Daily attendance processed successfully", result)
}

func (h *Handler) GetSummary(c *gin.Context) {
	employeeIDStr := c.Query("employeeId")
	startDateStr := c.Query("startDate")
	endDateStr := c.Query("endDate")

	empID, err := uuid.Parse(employeeIDStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid employeeId format")
		return
	}

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid startDate (use YYYY-MM-DD)")
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid endDate (use YYYY-MM-DD)")
		return
	}

	summary, err := h.service.GetMonthlySummary(c.Request.Context(), empID, startDate, endDate)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, "Attendance summary fetched successfully", summary)
}
