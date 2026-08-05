package auth

import (
	"github.com/tamim1715/novaerp/internal/app"
	"github.com/tamim1715/novaerp/internal/modules/user"
)

func NewModule(app *app.Application) *Handler {
	userRepo := user.NewRepository(app.DB)
	service := NewService(userRepo, app.Config.JWTSecret, app.Logger)
	return NewHandler(service)
}
