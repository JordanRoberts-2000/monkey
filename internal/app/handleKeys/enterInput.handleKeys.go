package handlekeys

import (
	"main/internal/state"

	"github.com/gdamore/tcell/v2"
)

func EnterInputKeys(appState *state.AppState, ev *tcell.EventKey) {
	switch ev.Key() {
	case tcell.KeyEnter:
		appState.CurrentScreen = state.MonkeyTyping
	case tcell.KeyBackspace, tcell.KeyBackspace2:
		if len(appState.UserInput) > 0 {
			appState.UserInput = appState.UserInput[:len(appState.UserInput)-1]
		}
	case tcell.KeyRune:
		if ev.Rune() == 'B' {
			appState.CurrentScreen = state.SelectKeys
		} else {
			appState.UserInput += string(ev.Rune())
		}
	}
}
