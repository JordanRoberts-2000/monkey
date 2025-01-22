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

		components.DrawHeader(app.Screen, "Infinite Monkey Theorem")

		switch app.State.CurrentScreen {
		case state.Intro:
			screens.IntroScreen(app.Screen)
		case state.SelectKeys:
			screens.SelectKeys(app.Screen)
		}
		// next screen after pressing any key

		app.Screen.Show()
		app.HandleKeys()
	}
}
