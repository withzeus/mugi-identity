package app

import "github.com/withzeus/mugi-identity/core"

type App struct {
	helper core.Helper
}

func NewApp() *App {
	h := core.Helper{}
	return &App{helper: h}
}

func (app *App) Run() {

}
