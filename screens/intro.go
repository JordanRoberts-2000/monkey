package screens

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"main/styles"
	"main/utils"

	"github.com/mitchellh/go-wordwrap"
)

func IntroScreen(height int, width int) {
	if height >= 9 {
		utils.ClearScreen()
	  utils.PrintHeader("Infinite Monkey Theorem")

		var multiplier float64
		if width > 60 {
			multiplier = 0.8
		} else {
			multiplier = 0.9
		}
		scaledWidth := int(float64(width) * multiplier)

		introduction := `If an immortal monkey were to hit keys on a typewriter at random for an infinite amount of time, it would type every possible finite sequence of text, including the complete works of Shakespeare, every book ever written, and even your own life story.`
		wrappedIntroduction := wordwrap.WrapString(introduction, uint(scaledWidth))
		

		bulletPointOne := "Given infinite time, any sequence of events with a non-zero probability will occur."
		wrappedBulletPointOne := wordwrap.WrapString(bulletPointOne, uint(scaledWidth))

		bulletPointTwo := "While the probability of typing any specific text (like a novel) is astronomically small, over infinite time, the sheer number of attempts ensures that it will eventually happen."
		wrappedBulletPointTwo := wordwrap.WrapString(bulletPointTwo, uint(scaledWidth))
		
		
		lines := strings.Split(wrappedIntroduction, "\n")
		for i, line := range lines {
			if i == 0 {
				styles.Highlighted.Print("\"")
				styles.Muted.Println(" " + line)
			} else if i == len(lines) - 1 {
				styles.Muted.Print("  " + line + " ")
				styles.Highlighted.Println("\"")
			}else {
				styles.Muted.Println("  " + line)
			}
		}

		fmt.Println()
		
		lines2 := strings.Split(wrappedBulletPointOne, "\n")
		for i, line := range lines2 {
			if i == 0 {
				styles.Highlighted.Print("○")
				styles.Muted.Println(" " + line)
			} else {
				styles.Muted.Println("  " + line)
			}
		}

		fmt.Println()
		
		lines3 := strings.Split(wrappedBulletPointTwo, "\n")
		for i, line := range lines3 {
			if i == 0 {
				styles.Highlighted.Print("○")
				styles.Muted.Println(" " + line)
			} else {
				styles.Muted.Println("  " + line)
			}
		}

		fmt.Println()
		fmt.Print("Press enter to continue ")
		reader := bufio.NewReader(os.Stdin)
		_, _ = reader.ReadByte()
	}
}