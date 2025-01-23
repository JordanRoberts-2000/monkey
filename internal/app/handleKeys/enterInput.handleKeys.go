package handlekeys

import (
	"main/internal/state"

	"github.com/gdamore/tcell/v2"
)

func EnterInputKeys(appState *state.AppState, ev *tcell.EventKey) {
	switch ev.Key() {
	case tcell.KeyRune:
		appState.UserInput += string(ev.Rune())
	case tcell.KeyBackspace, tcell.KeyBackspace2:
		if len(appState.UserInput) > 0 {
			appState.UserInput = appState.UserInput[:len(appState.UserInput)-1]
		}
	}
}
