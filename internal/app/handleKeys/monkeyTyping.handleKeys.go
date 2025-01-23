package handlekeys

import (
	"main/internal/state"

	"github.com/gdamore/tcell/v2"
)

func MonkeyTypingKeys(appState *state.AppState, ev *tcell.EventKey) {
	switch ev.Key() {
	case tcell.KeyRune:
		if ev.Rune() == 'B' {
			appState.CurrentScreen = state.EnterInput
			appState.UserInput = ""
			appState.InvalidMessage = ""
		}
	}
}
