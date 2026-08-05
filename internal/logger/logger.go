// Package logger provides structured logging initialization for the application.
package logger

import "go.uber.org/zap"

// NewLogger initializes a zap.Logger instance or returns an error.
func NewLogger() (*zap.Logger, error) {
	log, err := zap.NewProduction()
	if err != nil {
		// Fallback to development logger if production configuration fails
		devLog, devErr := zap.NewDevelopment()
		if devErr != nil {
			return nil, err
		}
		return devLog, nil
	}

	return log, nil
}
