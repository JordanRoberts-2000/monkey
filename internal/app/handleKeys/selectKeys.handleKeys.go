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
	case tcell.KeyEnter:
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
