package report

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/common/response"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetGeneralLedger(c *gin.Context) {
	var accountID *uuid.UUID
	if accStr := c.Query("accountId"); accStr != "" {
		if uid, err := uuid.Parse(accStr); err == nil {
			accountID = &uid
		}
	}

	startDate := time.Date(time.Now().Year(), 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Now().UTC()

	if s := c.Query("startDate"); s != "" {
		if t, err := time.Parse("2006-01-02", s); err == nil {
			startDate = t
		}
	}
	if e := c.Query("endDate"); e != "" {
		if t, err := time.Parse("2006-01-02", e); err == nil {
			endDate = t
		}
	}

	gl, err := h.service.GetGeneralLedger(c.Request.Context(), accountID, startDate, endDate)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, "General Ledger statement generated successfully", gl)
}

func (h *Handler) GetTrialBalance(c *gin.Context) {
	asOfDate := time.Now().UTC()
	if d := c.Query("asOfDate"); d != "" {
		if t, err := time.Parse("2006-01-02", d); err == nil {
			asOfDate = t
		}
	}

	tb, err := h.service.GetTrialBalance(c.Request.Context(), asOfDate)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, "Trial Balance report generated successfully", tb)
}

func (h *Handler) GetProfitAndLoss(c *gin.Context) {
	startDate := time.Date(time.Now().Year(), 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Now().UTC()

	if s := c.Query("startDate"); s != "" {
		if t, err := time.Parse("2006-01-02", s); err == nil {
			startDate = t
		}
	}
	if e := c.Query("endDate"); e != "" {
		if t, err := time.Parse("2006-01-02", e); err == nil {
			endDate = t
		}
	}

	pl, err := h.service.GetProfitAndLoss(c.Request.Context(), startDate, endDate)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, "Profit and Loss statement generated successfully", pl)
}

func (h *Handler) GetBalanceSheet(c *gin.Context) {
	asOfDate := time.Now().UTC()
	if d := c.Query("asOfDate"); d != "" {
		if t, err := time.Parse("2006-01-02", d); err == nil {
			asOfDate = t
		}
	}

	bs, err := h.service.GetBalanceSheet(c.Request.Context(), asOfDate)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, "Balance Sheet statement generated successfully", bs)
}
