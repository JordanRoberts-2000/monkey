package handlekeys

import (
	"main/internal/state"
)

func IntroScreenKeys(appState *state.AppState) {
	appState.CurrentScreen = state.SelectKeys
}
