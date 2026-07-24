package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/tamim1715/novaerp/internal/app"
	"github.com/tamim1715/novaerp/internal/common/middleware"
	"github.com/tamim1715/novaerp/internal/config"
	"github.com/tamim1715/novaerp/internal/database"
	"github.com/tamim1715/novaerp/internal/logger"
	"github.com/tamim1715/novaerp/internal/routes"
)

func Bootstrap() (*gin.Engine, *app.Application) {

	// Load environment
	config.LoadEnv()

	// Initialize logger
	log := logger.NewLogger()

	// Connect database
	database.Connect()

	// Create application context
	application := &app.Application{
		Config: &config.AppConfig,
		DB:     database.DB,
		Logger: log,
	}

	// Configure Gin
	gin.SetMode(application.Config.GinMode)

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(middleware.RequestID())

	// Register all routes
	routes.RegisterRoutes(router, application)

	return router, application
}
