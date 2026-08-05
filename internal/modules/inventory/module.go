package inventory

import "github.com/tamim1715/novaerp/internal/app"

func NewModule(app *app.Application) *Handler {
	repo := NewRepository(app.DB)
	service := NewService(repo, app.Logger)
	return NewHandler(service)
}
