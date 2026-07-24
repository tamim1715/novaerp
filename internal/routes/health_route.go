package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tamim1715/novaerp/internal/app"
	"github.com/tamim1715/novaerp/internal/handler"
)

func RegisterHealthRoutes(router *gin.Engine, application *app.Application) {

	healthHandler := handler.NewHealthHandler(application)

	router.GET("/", healthHandler.Running)
	router.GET("/health", healthHandler.Health)

}
