package screens

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/inancgumus/screen"
	"github.com/mitchellh/go-wordwrap"
)

func IntroScreen(height int, width int) {
	if height >= 9 {
		screen.Clear()
		screen.MoveTopLeft()
	
		display := color.New(color.FgWhite, color.BgBlack, color.Bold)
		boldText := color.New(color.Bold)
		display.Print(" Monkey ")

		text := `"If an immortal monkey were to hit keys on a typewriter at random for an infinite amount of time, it would almost surely type every possible finite sequence of text, including the complete works of Shakespeare, every book ever written, and even your own life story."`
		scaledWidth := int(float64(width) * 0.7)
		wrapped := wordwrap.WrapString(text, uint(scaledWidth))
		
		boldText.Println("  \"Infinite monkey theorem\"")
		fmt.Println()
		fmt.Println(wrapped)
		fmt.Println()
		fmt.Println("∘ Given infinite time, any sequence of events with a non-zero probability will occur.")
		fmt.Println()
		fmt.Println("∘ While the probability of typing any specific text (like a novel) is astronomically small, over \n  infinite time, the sheer number of attempts ensures that it will eventually happen.")
		fmt.Println()
		fmt.Print("Press enter to continue...")
		reader := bufio.NewReader(os.Stdin)
		_, _ = reader.ReadByte()
	}
}