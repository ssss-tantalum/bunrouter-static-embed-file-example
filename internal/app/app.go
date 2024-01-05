package app

import (
	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/bunrouterotel"
	"github.com/uptrace/bunrouter/extra/reqlog"
)

type App struct {
	app *App

	router *bunrouter.Router
}

func New() *App {
	app := &App{}
	return app
}

func (app *App) Router() *bunrouter.Router {
	app.router = bunrouter.New(
		bunrouter.WithMiddleware(bunrouterotel.NewMiddleware()),
		bunrouter.WithMiddleware(reqlog.NewMiddleware(
			reqlog.WithEnabled(true),
			reqlog.FromEnv(""),
		)),
	)

	return app.router
}
