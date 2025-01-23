package main

import (
	"log"
	"main/internal/app"
	"main/internal/components"
	"main/internal/screens"
	"main/internal/state"
)

func main() {
	app, err := app.Initialize()
	if err != nil {
		log.Fatalf("Failed to initialise app: %v", err)
	}
	defer app.Screen.Fini()

	for app.State.Running {
		app.Screen.Clear()

		switch app.State.CurrentScreen {
		case state.Intro:
			components.DrawHeader(app.Screen, "Infinite Monkey Theorem")
			screens.IntroScreen(app.Screen)
		case state.SelectKeys:
			components.DrawHeader(app.Screen, "Select Character Set:")
			screens.SelectKeys(app.Screen, int(app.State.CharSetIndex))
		case state.EnterInput:
			components.DrawHeader(app.Screen, "Enter Some Target Text:")
			screens.EnterInput(app.Screen)
		case state.MonkeyTyping:
			components.DrawHeader(app.Screen, "Monkey Is Typing!")
			screens.MonkeyTyping(app.Screen)
		}

		app.Screen.Show()
		app.HandleKeys()
	}
}
