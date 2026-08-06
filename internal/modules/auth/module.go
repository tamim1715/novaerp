package auth

import (
	"github.com/tamim1715/novaerp/internal/app"
	"github.com/tamim1715/novaerp/internal/modules/user"
)

func NewModule(app *app.Application, keyManager *KeyManager) *Handler {
	authRepo := NewRepository(app.DB)
	userRepo := user.NewRepository(app.DB)
	service := NewService(authRepo, userRepo, keyManager, app.Logger)
	return NewHandler(service)
}
