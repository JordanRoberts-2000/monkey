package handlekeys

import (
	"main/internal/state"

	"github.com/gdamore/tcell/v2"
)

func MonkeyTypingKeys(appState *state.AppState, ev *tcell.EventKey) {
	switch ev.Key() {
	case tcell.KeyLeft:
		appState.CurrentScreen = state.EnterInput
		if appState.UserInput != "" {
			appState.UserInput = ""
		}
		if appState.InvalidMessage != "" {
			appState.InvalidMessage = ""
		}
	}
}
