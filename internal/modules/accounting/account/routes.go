package account

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
	router.POST("", handler.CreateAccountDoc)
	router.GET("", handler.FindAllAccountsDoc)
	router.GET("/tree", handler.GetAccountTreeDoc)
	router.POST("/seed", handler.SeedAccountsDoc)
	router.GET("/:id", handler.FindAccountByIDDoc)
	router.PUT("/:id", handler.UpdateAccountDoc)
	router.DELETE("/:id", handler.DeleteAccountDoc)
}
