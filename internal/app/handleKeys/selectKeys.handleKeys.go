package handlekeys

import (
	"main/internal/state"

	"github.com/gdamore/tcell/v2"
)

func SelectScreenKeys(appState *state.AppState, ev *tcell.EventKey) {
	switch ev.Key() {
	case tcell.KeyUp:
		appState.CharSetIndex = state.CharSetId((int(appState.CharSetIndex) - 1 + 4) % 4)
	case tcell.KeyDown:
		appState.CharSetIndex = state.CharSetId((int(appState.CharSetIndex) + 1) % 4)
	case tcell.KeyLeft:
		appState.CurrentScreen = state.Intro
		if appState.UserInput != "" {
			appState.UserInput = ""
		}
		if appState.InvalidMessage != "" {
			appState.InvalidMessage = ""
		}
	case tcell.KeyEnter:
		appState.CurrentScreen = state.EnterInput
	case tcell.KeyRight:
		appState.CurrentScreen = state.EnterInput
	default:
		switch ev.Rune() {
		case 'w':
			appState.CharSetIndex = state.CharSetId((int(appState.CharSetIndex) - 1 + 4) % 4)
		case 's':
			appState.CharSetIndex = state.CharSetId((int(appState.CharSetIndex) + 1) % 4)
		}
	}
}
