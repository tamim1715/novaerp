package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tamim1715/novaerp/internal/app"
	"github.com/tamim1715/novaerp/internal/modules/department"
)

func RegisterRoutes(router *gin.Engine, application *app.Application) {

	api := router.Group("/api/v1")

	// Health Route
	RegisterHealthRoutes(router, application)

	// Department Module
	departmentHandler := department.NewModule(application)
	department.RegisterRoutes(api.Group("/departments"), departmentHandler)
}
