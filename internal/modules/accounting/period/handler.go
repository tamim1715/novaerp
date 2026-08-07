package period

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamim1715/novaerp/internal/common/response"
	"gorm.io/gorm"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateFiscalYear(c *gin.Context) {
	var req CreateFiscalYearRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	fy, err := h.service.CreateFiscalYear(c.Request.Context(), req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Created(c, "Fiscal year created successfully", ToFiscalYearResponse(fy))
}

func (h *Handler) FindAllFiscalYears(c *gin.Context) {
	fys, err := h.service.FindAllFiscalYears(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, "Fiscal years fetched successfully", ToFiscalYearResponseList(fys))
}

func (h *Handler) FindFiscalYearByID(c *gin.Context) {
	id := c.Param("id")
	fy, err := h.service.FindFiscalYearByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "Fiscal year not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, "Fiscal year fetched successfully", ToFiscalYearResponse(fy))
}

func (h *Handler) CloseFiscalYear(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.CloseFiscalYear(c.Request.Context(), id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "Fiscal year not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, "Fiscal year and all its periods closed successfully", nil)
}

func (h *Handler) FindPeriodByID(c *gin.Context) {
	id := c.Param("id")
	p, err := h.service.FindPeriodByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "Accounting period not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, "Accounting period fetched successfully", ToAccountingPeriodResponse(p))
}

func (h *Handler) ClosePeriod(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.ClosePeriod(c.Request.Context(), id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "Accounting period not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, "Accounting period closed successfully", nil)
}

func (h *Handler) ReopenPeriod(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.ReopenPeriod(c.Request.Context(), id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "Accounting period not found")
			return
		}
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, "Accounting period reopened successfully", nil)
}
