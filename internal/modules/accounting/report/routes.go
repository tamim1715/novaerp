package report

import (
	"github.com/gin-gonic/gin"
	"github.com/tamim1715/novaerp/internal/app"
	"github.com/tamim1715/novaerp/internal/modules/accounting/account"
	"github.com/tamim1715/novaerp/internal/modules/accounting/journal"
)

func NewModule(app *app.Application) *Handler {
	accountRepo := account.NewRepository(app.DB)
	journalRepo := journal.NewRepository(app.DB)

	service := NewService(accountRepo, journalRepo, app.Logger)
	return NewHandler(service)
}

func RegisterRoutes(router *gin.RouterGroup, handler *Handler) {
	reportGroup := router.Group("/reports")
	{
		reportGroup.GET("/general-ledger", handler.GetGeneralLedgerDoc)
		reportGroup.GET("/trial-balance", handler.GetTrialBalanceDoc)
		reportGroup.GET("/profit-loss", handler.GetProfitAndLossDoc)
		reportGroup.GET("/balance-sheet", handler.GetBalanceSheetDoc)
	}
}
