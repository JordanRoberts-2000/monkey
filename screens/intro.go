package screens

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/inancgumus/screen"
)

func IntroScreen() {
	screen.Clear()
	screen.MoveTopLeft()

	display := color.New(color.FgWhite, color.BgGreen, color.Bold)
	boldText := color.New(color.Bold)
	display.Print(" Monkey ")

	boldText.Println("  \"Infinite monkey theorem\"")
	fmt.Println()
	fmt.Println("\"If an immortal monkey were to hit keys on a typewriter at random for an infinite amount of time, \n it would almost surely type every possible finite sequence of text, including the complete works \n of Shakespeare, every book ever written, and even your own life story.\"")
	fmt.Println()
	fmt.Println("∘ Given infinite time, any sequence of events with a non-zero probability will occur.")
	fmt.Println()
	fmt.Println("∘ While the probability of typing any specific text (like a novel) is astronomically small, over \n  infinite time, the sheer number of attempts ensures that it will eventually happen.")
	fmt.Println()
	fmt.Print("Press enter to continue...")
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadByte()
}