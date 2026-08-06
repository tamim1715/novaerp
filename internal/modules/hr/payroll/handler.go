package payroll

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

func (h *Handler) CreatePeriod(c *gin.Context) {
	var req CreatePayrollPeriodRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	period, err := h.service.CreatePeriod(c.Request.Context(), req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Created(c, "Payroll period created successfully", ToPayrollPeriodResponse(period))
}

func (h *Handler) ProcessPayroll(c *gin.Context) {
	id := c.Param("id")
	var req ProcessPayrollRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	period, err := h.service.ProcessPayroll(c.Request.Context(), id, req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, "Payroll processed successfully", ToPayrollPeriodResponse(period))
}

func (h *Handler) FindAllPeriods(c *gin.Context) {
	var req pagination.PageRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	list, total, err := h.service.FindAllPeriods(c.Request.Context(), req)
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
		Data:       ToPayrollPeriodResponseList(list),
	}

	response.Success(c, "Payroll periods fetched successfully", pageResp)
}

func (h *Handler) FindPeriodByID(c *gin.Context) {
	id := c.Param("id")
	period, err := h.service.FindPeriodByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "Payroll period not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, "Payroll period fetched successfully", ToPayrollPeriodResponse(period))
}

func (h *Handler) GetPayslips(c *gin.Context) {
	id := c.Param("id")
	payslips, err := h.service.GetPayslipsByPeriodID(c.Request.Context(), id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, "Payslips fetched successfully", ToPayslipResponseList(payslips))
}

func (h *Handler) MarkPaid(c *gin.Context) {
	id := c.Param("id")
	period, err := h.service.MarkPaid(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "Payroll period not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, "Payroll marked as paid successfully", ToPayrollPeriodResponse(period))
}
