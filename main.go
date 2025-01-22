package main

import (
	"log"
	"main/styles"
	"main/utils"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/mitchellh/go-wordwrap"
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

		width, _ := screen.Size()
		var multiplier float64
		if width > 60 {
			multiplier = 0.8
		} else {
			multiplier = 0.98
		}
		scaledWidth := int(float64(width) * multiplier)

		introduction := `If an immortal monkey were to hit keys on a typewriter at random for an infinite amount of time, it would type every possible finite sequence of text, including the complete works of Shakespeare, every book ever written, and even your own life story.`
		wrappedIntroduction := wordwrap.WrapString(introduction, uint(scaledWidth))

		lines := strings.Split(wrappedIntroduction, "\n")

		y := 2 
		for _, line := range lines {
			for x, ch := range line {
				screen.SetContent(x, y, ch, nil, tcell.StyleDefault)
			}
			y++ // Move to the next line
		}

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
