package app

import (
	handlekeys "main/internal/app/handleKeys"
	"main/internal/state"

	"github.com/gdamore/tcell/v2"
)

type App struct {
	Screen tcell.Screen
	State  *state.AppState
}

func Initialize() (*App, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}
	if err := screen.Init(); err != nil {
		return nil, err
	}
	return &App{
		Screen: screen,
		State:  state.Initialize(),
	}, nil
}

func (app *App) HandleKeys() {
	handlekeys.HandleKeys(app.Screen, app.State)
}
