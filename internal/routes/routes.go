package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tamim1715/novaerp/internal/app"
	"github.com/tamim1715/novaerp/internal/modules/auth"
	"github.com/tamim1715/novaerp/internal/modules/department"
	"github.com/tamim1715/novaerp/internal/modules/employee"
	"github.com/tamim1715/novaerp/internal/modules/inventory"
	"github.com/tamim1715/novaerp/internal/modules/user"
)

func RegisterRoutes(router *gin.Engine, application *app.Application) {

	api := router.Group("/api/v1")

	// Health Route
	RegisterHealthRoutes(router, application)

	// Auth Module
	authHandler := auth.NewModule(application)
	auth.RegisterRoutes(api.Group("/auth"), authHandler)

	// Department Module
	departmentHandler := department.NewModule(application)
	department.RegisterRoutes(api.Group("/departments"), departmentHandler)

	// User Module
	userHandler := user.NewModule(application)
	user.RegisterRoutes(api.Group("/users"), userHandler)

	// Employee module
	employeeHandler := employee.NewModule(application)
	employee.RegisterRoutes(api.Group("/employees"), employeeHandler)

	// Inventory module
	inventoryHandler := inventory.NewModule(application)
	inventory.RegisterRoutes(api, inventoryHandler)
}
