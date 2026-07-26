package user

import "github.com/tamim1715/novaerp/internal/app"

func NewModule(app *app.Application) *Handler {

	repository := NewRepository(app.DB)

	service := NewService(
		repository,
		app.Logger,
	)

	return NewHandler(service)
}
