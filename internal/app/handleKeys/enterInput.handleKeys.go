package handlekeys

import (
	"main/internal/state"
	"main/internal/utils"

	"github.com/gdamore/tcell/v2"
)

func EnterInputKeys(appState *state.AppState, ev *tcell.EventKey) {
	switch ev.Key() {
	case tcell.KeyEnter:
		if utils.ValidateInput(appState) {
			appState.CurrentScreen = state.MonkeyTyping
		}
	case tcell.KeyRight:
		if utils.ValidateInput(appState) {
			appState.CurrentScreen = state.MonkeyTyping
		}
	case tcell.KeyBackspace, tcell.KeyBackspace2:
		if len(appState.UserInput) > 0 {
			appState.UserInput = appState.UserInput[:len(appState.UserInput)-1]
		}
		if appState.InvalidMessage != "" {
			appState.InvalidMessage = ""
		}
	case tcell.KeyLeft:
		appState.CurrentScreen = state.SelectKeys
		if appState.UserInput != "" {
			appState.UserInput = ""
		}
		if appState.InvalidMessage != "" {
			appState.InvalidMessage = ""
		}

	case tcell.KeyRune:
		if len(appState.UserInput) < 50 {
			appState.UserInput += string(ev.Rune())
		}
	}
}
