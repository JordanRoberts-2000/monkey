package main

import (
	"log"
	"main/styles"
	"main/utils"

	"github.com/gdamore/tcell/v2"
)


func main() {
	// Initialize tcell screen
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("Failed to create tcell screen: %v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("Failed to initialize tcell screen: %v", err)
	}
	defer screen.Fini() 

	// loop
	running := true
	for running {
		screen.Clear()

		utils.DrawText(screen, 0, 0, " Monkey ", styles.Display)
		
		utils.DrawText(screen, 10, 0, "Infinite Monkey Theorem", styles.Default.Bold(true))

		screen.Show()
		// getting the event
		ev := screen.PollEvent()
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
