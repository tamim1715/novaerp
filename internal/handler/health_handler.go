package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tamim1715/novaerp/internal/app"
	"github.com/tamim1715/novaerp/internal/common/response"
	"github.com/tamim1715/novaerp/internal/config"
)

type HealthHandler struct {
	app *app.Application
}

func NewHealthHandler(app *app.Application) *HealthHandler {
	return &HealthHandler{
		app: app,
	}
}

func (h *HealthHandler) Health(c *gin.Context) {
	response.Success(c, "Application is healthy", gin.H{
		"application": config.AppConfig.AppName,
		"status":      "UP",
	})

}

func (h *HealthHandler) Running(c *gin.Context) {
	response.Success(c, "NovaERP Running", gin.H{
		"application": config.AppConfig.AppName,
		"status":      "running",
	})
}
