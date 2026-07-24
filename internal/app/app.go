package app

import (
	"github.com/tamim1715/novaerp/internal/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Application struct {
	Config *config.Config
	DB     *gorm.DB
	Logger *zap.Logger
}
