package journal

import (
	"github.com/gin-gonic/gin"
	"github.com/tamim1715/novaerp/internal/app"
	"github.com/tamim1715/novaerp/internal/modules/accounting/account"
	"github.com/tamim1715/novaerp/internal/modules/accounting/period"
)

func NewModule(app *app.Application) *Handler {
	repo := NewRepository(app.DB)
	accountRepo := account.NewRepository(app.DB)
	periodRepo := period.NewRepository(app.DB)

	service := NewService(repo, accountRepo, periodRepo, app.Logger)
	return NewHandler(service)
}

func RegisterRoutes(router *gin.RouterGroup, handler *Handler) {
	router.POST("", handler.CreateJournalEntryDoc)
	router.GET("", handler.FindAllJournalEntriesDoc)
	router.GET("/:id", handler.FindJournalEntryByIDDoc)
	router.POST("/:id/post", handler.PostJournalEntryDoc)
	router.POST("/:id/void", handler.VoidJournalEntryDoc)
}
