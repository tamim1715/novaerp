// @title Nova ERP API
// @version 1.0
// @description Nova ERP REST API
// @BasePath /api/v1
// @host localhost:8080

package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/tamim1715/novaerp/docs"
	"github.com/tamim1715/novaerp/internal/bootstrap"
	"go.uber.org/zap"
)

func main() {
	router, application, err := bootstrap.Bootstrap()
	if err != nil {
		log.Fatalf("failed to bootstrap application: %v", err)
	}

	defer application.Logger.Sync()

	// Register Swagger route
	router.GET(
		"/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler),
	)

	serverAddr := ":" + application.Config.Port
	srv := &http.Server{
		Addr:         serverAddr,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start HTTP server in a separate goroutine
	go func() {
		application.Logger.Info("server starting",
			zap.String("app", application.Config.AppName),
			zap.String("port", application.Config.Port),
		)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			application.Logger.Fatal("server error", zap.Error(err))
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	application.Logger.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		application.Logger.Error("server forced to shutdown", zap.Error(err))
	} else {
		application.Logger.Info("server exited gracefully")
	}

	// Close database connection pool gracefully
	if sqlDB, err := application.DB.DB(); err == nil {
		application.Logger.Info("closing database connection pool...")
		if err := sqlDB.Close(); err != nil {
			application.Logger.Error("failed to close database connection pool", zap.Error(err))
		} else {
			application.Logger.Info("database connection pool closed gracefully")
		}
	}
}
