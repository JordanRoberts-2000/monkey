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
			components.DrawHeader(app.Screen, "Select Allowed Keys:")
			screens.SelectKeys(app.Screen)
		}

		app.Screen.Show()
		app.HandleKeys()
	}
}
