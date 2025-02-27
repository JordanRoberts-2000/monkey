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
			components.DrawHeader(app.Screen, "Infinite Monkey Theorem", 1)
			screens.IntroScreen(app.Screen)
		case state.SelectKeys:
			components.DrawHeader(app.Screen, "Select Character Set:", 2)
			screens.SelectKeys(app.Screen, int(app.State.CharSetIndex))
		case state.EnterInput:
			components.DrawHeader(app.Screen, "Enter Some Target Text:", 3)
			screens.EnterInput(app.Screen, app.State.UserInput, app.State.InvalidMessage, int(app.State.CharSetIndex))
		case state.MonkeyTyping:
			components.DrawHeader(app.Screen, "Monkey Is Typing!", 4)
			screens.MonkeyTyping(app.Screen, app.State.UserInput)
		}

		app.Screen.Show()
		app.HandleKeys()
	}
}
