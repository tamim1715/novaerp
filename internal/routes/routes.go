package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tamim1715/novaerp/internal/app"
)

func RegisterRoutes(router *gin.Engine, application *app.Application) {

	RegisterHealthRoutes(router, application)
}
