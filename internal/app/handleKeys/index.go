package handlekeys

import (
	"main/internal/state"

	"github.com/gdamore/tcell/v2"
)

func HandleKeys(screen tcell.Screen, appState *state.AppState) {
	event := screen.PollEvent()
	if event == nil {
		return // No event to handle
	}

	// Ensure the event is a key event
	keyEvent, ok := event.(*tcell.EventKey)
	if !ok {
		return
	}

	GlobalKeys(appState, keyEvent)
	switch appState.CurrentScreen {
	case state.Intro:
		IntroScreenKeys(appState)
	case state.SelectKeys:
		SelectScreenKeys(appState, keyEvent)
	}

}
