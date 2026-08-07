package report

import (
	"github.com/gin-gonic/gin"
	"github.com/tamim1715/novaerp/internal/common/response"
)

var (
	_ response.APIResponse
)

// GetGeneralLedgerDoc generates running General Ledger transaction report
// @Summary General Ledger Report
// @Description Generates detailed running ledger with opening balance, posted debits/credits, and closing balances
// @Tags Accounting - Financial Reports & Statements
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param accountId query string false "Filter by Account UUID (optional, defaults to all active accounts)"
// @Param startDate query string false "Start date filter (YYYY-MM-DD)"
// @Param endDate query string false "End date filter (YYYY-MM-DD)"
// @Success 200 {object} response.APIResponse{data=GeneralLedgerResponse} "General Ledger statement generated"
// @Failure 401 {object} response.APIResponse "Unauthorized"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /accounting/reports/general-ledger [get]
func (h *Handler) GetGeneralLedgerDoc(c *gin.Context) {
	h.GetGeneralLedger(c)
}

// GetTrialBalanceDoc generates Trial Balance statement
// @Summary Trial Balance Statement
// @Description Verifies ledger debit and credit turnover equality and net balances across all accounts
// @Tags Accounting - Financial Reports & Statements
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param asOfDate query string false "As of date filter (YYYY-MM-DD, defaults to today)"
// @Success 200 {object} response.APIResponse{data=TrialBalanceResponse} "Trial Balance statement generated"
// @Failure 401 {object} response.APIResponse "Unauthorized"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /accounting/reports/trial-balance [get]
func (h *Handler) GetTrialBalanceDoc(c *gin.Context) {
	h.GetTrialBalance(c)
}

// GetProfitAndLossDoc generates Income Statement / P&L
// @Summary Profit and Loss (Income Statement)
// @Description Computes Total Revenues, COGS, Gross Profit, Operating Expenses, and Net Income for a period
// @Tags Accounting - Financial Reports & Statements
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param startDate query string false "Start date (YYYY-MM-DD, defaults to start of current year)"
// @Param endDate query string false "End date (YYYY-MM-DD, defaults to today)"
// @Success 200 {object} response.APIResponse{data=ProfitAndLossResponse} "Profit and Loss statement generated"
// @Failure 401 {object} response.APIResponse "Unauthorized"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /accounting/reports/profit-loss [get]
func (h *Handler) GetProfitAndLossDoc(c *gin.Context) {
	h.GetProfitAndLoss(c)
}

// GetBalanceSheetDoc generates Balance Sheet statement
// @Summary Balance Sheet Statement
// @Description Evaluates Assets = Liabilities + Equity (including current period earnings) as of a specific date
// @Tags Accounting - Financial Reports & Statements
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param asOfDate query string false "As of date (YYYY-MM-DD, defaults to today)"
// @Success 200 {object} response.APIResponse{data=BalanceSheetResponse} "Balance sheet statement generated"
// @Failure 401 {object} response.APIResponse "Unauthorized"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /accounting/reports/balance-sheet [get]
func (h *Handler) GetBalanceSheetDoc(c *gin.Context) {
	h.GetBalanceSheet(c)
}
