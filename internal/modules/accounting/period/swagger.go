package period

import (
	"github.com/gin-gonic/gin"
	"github.com/tamim1715/novaerp/internal/common/response"
)

var (
	_ response.APIResponse
)

// CreateFiscalYearDoc creates a new fiscal year and optionally auto-generates 12 monthly accounting periods
// @Summary Create Fiscal Year
// @Description Register a new fiscal year with automated 12 monthly accounting sub-periods
// @Tags Accounting - Periods & Fiscal Years
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body CreateFiscalYearRequest true "Fiscal Year Payload"
// @Success 201 {object} response.APIResponse{data=FiscalYearResponse} "Fiscal year created successfully"
// @Failure 400 {object} response.APIResponse "Bad request"
// @Failure 401 {object} response.APIResponse "Unauthorized"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /accounting/fiscal-years [post]
func (h *Handler) CreateFiscalYearDoc(c *gin.Context) {
	h.CreateFiscalYear(c)
}

// FindAllFiscalYearsDoc retrieves all fiscal years with their monthly periods
// @Summary List Fiscal Years
// @Description Retrieve list of all fiscal years and child monthly periods
// @Tags Accounting - Periods & Fiscal Years
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.APIResponse{data=[]FiscalYearResponse} "Fiscal years fetched successfully"
// @Failure 401 {object} response.APIResponse "Unauthorized"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /accounting/fiscal-years [get]
func (h *Handler) FindAllFiscalYearsDoc(c *gin.Context) {
	h.FindAllFiscalYears(c)
}

// FindFiscalYearByIDDoc retrieves a single fiscal year
// @Summary Get Fiscal Year by ID
// @Description Retrieve a single fiscal year by its UUID
// @Tags Accounting - Periods & Fiscal Years
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Fiscal Year UUID"
// @Success 200 {object} response.APIResponse{data=FiscalYearResponse} "Fiscal year fetched successfully"
// @Failure 404 {object} response.APIResponse "Fiscal year not found"
// @Failure 401 {object} response.APIResponse "Unauthorized"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /accounting/fiscal-years/{id} [get]
func (h *Handler) FindFiscalYearByIDDoc(c *gin.Context) {
	h.FindFiscalYearByID(c)
}

// CloseFiscalYearDoc closes a fiscal year and all associated periods
// @Summary Close Fiscal Year
// @Description Lock a fiscal year and freeze all child accounting periods from further postings
// @Tags Accounting - Periods & Fiscal Years
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Fiscal Year UUID"
// @Success 200 {object} response.APIResponse "Fiscal year closed successfully"
// @Failure 404 {object} response.APIResponse "Fiscal year not found"
// @Failure 401 {object} response.APIResponse "Unauthorized"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /accounting/fiscal-years/{id}/close [post]
func (h *Handler) CloseFiscalYearDoc(c *gin.Context) {
	h.CloseFiscalYear(c)
}

// FindPeriodByIDDoc retrieves an accounting period
// @Summary Get Accounting Period by ID
// @Description Retrieve a single accounting period by UUID
// @Tags Accounting - Periods & Fiscal Years
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Accounting Period UUID"
// @Success 200 {object} response.APIResponse{data=AccountingPeriodResponse} "Accounting period fetched successfully"
// @Failure 404 {object} response.APIResponse "Period not found"
// @Failure 401 {object} response.APIResponse "Unauthorized"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /accounting/periods/{id} [get]
func (h *Handler) FindPeriodByIDDoc(c *gin.Context) {
	h.FindPeriodByID(c)
}

// ClosePeriodDoc closes a monthly accounting period
// @Summary Close Accounting Period
// @Description Freeze a monthly accounting period to disallow backdated postings
// @Tags Accounting - Periods & Fiscal Years
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Accounting Period UUID"
// @Success 200 {object} response.APIResponse "Period closed successfully"
// @Failure 404 {object} response.APIResponse "Period not found"
// @Failure 401 {object} response.APIResponse "Unauthorized"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /accounting/periods/{id}/close [post]
func (h *Handler) ClosePeriodDoc(c *gin.Context) {
	h.ClosePeriod(c)
}

// ReopenPeriodDoc reopens a closed accounting period
// @Summary Reopen Accounting Period
// @Description Reopen a monthly period for authorized adjustments (if fiscal year is still open)
// @Tags Accounting - Periods & Fiscal Years
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Accounting Period UUID"
// @Success 200 {object} response.APIResponse "Period reopened successfully"
// @Failure 400 {object} response.APIResponse "Cannot reopen period of closed fiscal year"
// @Failure 404 {object} response.APIResponse "Period not found"
// @Failure 401 {object} response.APIResponse "Unauthorized"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /accounting/periods/{id}/reopen [post]
func (h *Handler) ReopenPeriodDoc(c *gin.Context) {
	h.ReopenPeriod(c)
}
