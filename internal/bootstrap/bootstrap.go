// Package bootstrap initializes application dependencies, database migrations, and routing.
package bootstrap

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tamim1715/novaerp/internal/app"
	"github.com/tamim1715/novaerp/internal/common/middleware"
	"github.com/tamim1715/novaerp/internal/config"
	"github.com/tamim1715/novaerp/internal/database"
	"github.com/tamim1715/novaerp/internal/logger"
	"github.com/tamim1715/novaerp/internal/modules/department"
	"github.com/tamim1715/novaerp/internal/modules/employee"
	"github.com/tamim1715/novaerp/internal/modules/user"
	"github.com/tamim1715/novaerp/internal/routes"
	"go.uber.org/zap"
)

// Bootstrap loads configuration, sets up logging, connects to the database, runs migrations, and configures Gin routes.
func Bootstrap() (*gin.Engine, *app.Application, error) {

	// Load environment
	cfg := config.LoadEnv()

	// Initialize logger
	log, err := logger.NewLogger()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to initialize logger: %w", err)
	}

	// Connect database
	db, err := database.Connect(cfg)
	if err != nil {
		log.Error("database connection error", zap.Error(err))
		return nil, nil, fmt.Errorf("database connection failed: %w", err)
	}

	// Create application context
	application := &app.Application{
		Config: cfg,
		DB:     db,
		Logger: log,
	}

	// Migrate the database
	if err := database.AutoMigrate(
		application.DB,
		&department.Department{},
		&user.User{},
		&employee.Employee{},
	); err != nil {
		log.Error("failed to migrate database", zap.Error(err))
		return nil, nil, fmt.Errorf("database migration failed: %w", err)
	}

	// Configure Gin
	gin.SetMode(application.Config.GinMode)

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(middleware.RequestID())

	// Register all routes
	routes.RegisterRoutes(router, application)

	return router, application, nil
}
