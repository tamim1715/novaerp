package period

import (
	"github.com/gin-gonic/gin"
	"github.com/tamim1715/novaerp/internal/app"
)

func NewModule(app *app.Application) *Handler {
	repo := NewRepository(app.DB)
	service := NewService(repo, app.Logger)
	return NewHandler(service)
}

func RegisterRoutes(router *gin.RouterGroup, handler *Handler) {
	// Fiscal Years
	fyGroup := router.Group("/fiscal-years")
	{
		fyGroup.POST("", handler.CreateFiscalYearDoc)
		fyGroup.GET("", handler.FindAllFiscalYearsDoc)
		fyGroup.GET("/:id", handler.FindFiscalYearByIDDoc)
		fyGroup.POST("/:id/close", handler.CloseFiscalYearDoc)
	}

	// Accounting Periods
	periodGroup := router.Group("/periods")
	{
		periodGroup.GET("/:id", handler.FindPeriodByIDDoc)
		periodGroup.POST("/:id/close", handler.ClosePeriodDoc)
		periodGroup.POST("/:id/reopen", handler.ReopenPeriodDoc)
	}
}
