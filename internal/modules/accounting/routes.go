package accounting

import (
	"github.com/gin-gonic/gin"
	"github.com/tamim1715/novaerp/internal/app"
	"github.com/tamim1715/novaerp/internal/modules/accounting/account"
	"github.com/tamim1715/novaerp/internal/modules/accounting/journal"
	"github.com/tamim1715/novaerp/internal/modules/accounting/period"
	"github.com/tamim1715/novaerp/internal/modules/accounting/report"
)

// RegisterRoutes aggregates and exposes all Accounting & Financial Management submodules.
func RegisterRoutes(protected *gin.RouterGroup, application *app.Application) {
	accountingGroup := protected.Group("/accounting")

	// 1. Chart of Accounts (COA)
	accountHandler := account.NewModule(application)
	account.RegisterRoutes(accountingGroup.Group("/accounts"), accountHandler)

	// 2. Fiscal Years & Accounting Periods
	periodHandler := period.NewModule(application)
	period.RegisterRoutes(accountingGroup, periodHandler)

	// 3. Journal Entries & General Ledger Engine
	journalHandler := journal.NewModule(application)
	journal.RegisterRoutes(accountingGroup.Group("/journals"), journalHandler)

	// 4. Financial Reports & Statements (General Ledger, Trial Balance, P&L, Balance Sheet)
	reportHandler := report.NewModule(application)
	report.RegisterRoutes(accountingGroup, reportHandler)
}
