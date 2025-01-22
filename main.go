package main

import (
	"log"
	"main/app"
	"main/components"
	"main/screens"

	"github.com/gdamore/tcell/v2"
)


func main() {
	app, err := app.Initialize()
	if err != nil {
		log.Fatalf("Failed to initialise app: %v", err)
	}
	defer app.Screen.Fini() 

	// loop
	running := true
	for running {
		app.Screen.Clear()

		components.DrawHeader(app.Screen, "Infinite Monkey Theorem")
		screens.IntroScreen(app.Screen)
		// next screen after pressing any key

		app.Screen.Show()
		// getting the event
		ev := app.Screen.PollEvent()
		// getting the event type
		switch ev := ev.(type) {
		case *tcell.EventKey:
			// getting the event key
			switch ev.Key() {
				case tcell.KeyEscape:
					running = false
				case tcell.KeyCtrlC:
					running = false
				case tcell.KeyCtrlQ:
					running = false
				case tcell.KeyCtrlX:
					running = false
				case tcell.KeyCtrlD:
					running = false
				default:
					switch ev.Rune() {
					case 'q':
						running = false
					}
				}
	}
}
}
