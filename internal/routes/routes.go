package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tamim1715/novaerp/internal/app"
	"github.com/tamim1715/novaerp/internal/common/middleware"
	"github.com/tamim1715/novaerp/internal/modules/auth"
	"github.com/tamim1715/novaerp/internal/modules/department"
	"github.com/tamim1715/novaerp/internal/modules/employee"
	"github.com/tamim1715/novaerp/internal/modules/hr"
	"github.com/tamim1715/novaerp/internal/modules/inventory"
	"github.com/tamim1715/novaerp/internal/modules/user"
)

func RegisterRoutes(router *gin.Engine, application *app.Application, keyManager *auth.KeyManager) {

	// Apply CORS Global Middleware
	router.Use(middleware.CORS(application.Config.CORSAllowedOrigins))

	// Base API Group
	api := router.Group("/api/v1")

	// Unprotected Routes
	RegisterHealthRoutes(router, application)

	// Auth Module (Public Auth Endpoints)
	authHandler := auth.NewModule(application, keyManager)
	auth.RegisterRoutes(api.Group("/auth"), authHandler)

	// User Module
	userHandler := user.NewModule(application)
	user.RegisterRoutes(api.Group("/users"), userHandler)

	// Protected Routes Group (Secured with RS256 Bearer JWT Auth Middleware)
	protected := api.Group("")
	protected.Use(auth.AuthMiddleware(keyManager.PublicKey))
	{
		// Department Module
		departmentHandler := department.NewModule(application)
		department.RegisterRoutes(protected.Group("/departments"), departmentHandler)

		// Employee module
		employeeHandler := employee.NewModule(application)
		employee.RegisterRoutes(protected.Group("/employees"), employeeHandler)

		// Inventory module
		inventory.RegisterRoutes(protected, application)

		// HR & Payroll module
		hr.RegisterRoutes(protected, application)
	}
}
