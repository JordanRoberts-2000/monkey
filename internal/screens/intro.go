package screens

import (
	"main/internal/styles"
	"main/internal/utils"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/mitchellh/go-wordwrap"
)

func IntroScreen(screen tcell.Screen) {
	width, height := screen.Size()
	var multiplier float64
	if width > 60 {
		multiplier = 0.8
	} else {
		multiplier = 0.98
	}
	scaledWidth := int(float64(width) * multiplier)

	introduction := `If an immortal monkey were to hit keys on a typewriter at random for an infinite amount of time, it would type every possible finite sequence of text, including the complete works of Shakespeare, every book ever written, and even your own life story.`
	wrappedIntroduction := wordwrap.WrapString(introduction, uint(scaledWidth))
	introductionLines := strings.Split(wrappedIntroduction, "\n")

	bulletOne := "Given infinite time, any sequence of events with a non-zero probability will occur."
	wrappedBulletOne := wordwrap.WrapString(bulletOne, uint(scaledWidth))
	bulletOneLines := strings.Split(wrappedBulletOne, "\n")

	bulletTwo := "While the probability of typing any specific text (like a novel) is astronomically small, over infinite time, the sheer number of attempts ensures that it will eventually happen."
	wrappedPointTwo := wordwrap.WrapString(bulletTwo, uint(scaledWidth))
	bulletTwoLines := strings.Split(wrappedPointTwo, "\n")

	y := 2
	for i, line := range introductionLines {
		if i == 0 {
			screen.SetContent(0, y, '"', nil, styles.Highlighted)
			for x, ch := range line {
				screen.SetContent(x+2, y, ch, nil, styles.Muted)
			}
		} else if i == len(introductionLines)-1 {
			for x, ch := range line {
				screen.SetContent(x+2, y, ch, nil, styles.Muted)
			}
			screen.SetContent(len(line)+3, y, '"', nil, styles.Highlighted)
		} else {
			for x, ch := range line {
				screen.SetContent(x+2, y, ch, nil, styles.Muted)
			}
		}
		y++
	}

	renderedHeight := len(bulletTwoLines) + len(bulletOneLines) + len(introductionLines) + 7 // header + footer + gap

	if height >= renderedHeight {
		y++

		for i, line := range bulletOneLines {
			if i == 0 {
				screen.SetContent(0, y, '○', nil, styles.Highlighted)
				for x, ch := range line {
					screen.SetContent(x+2, y, ch, nil, styles.Muted)
				}
			} else {
				for x, ch := range line {
					screen.SetContent(x+2, y, ch, nil, styles.Muted)
				}
			}
			y++
		}

		y++

		for i, line := range bulletTwoLines {
			if i == 0 {
				screen.SetContent(0, y, '○', nil, styles.Highlighted)
				for x, ch := range line {
					screen.SetContent(x+2, y, ch, nil, styles.Muted)
				}
			} else {
				for x, ch := range line {
					screen.SetContent(x+2, y, ch, nil, styles.Muted)
				}
			}
			y++
		}
	}

	y++

	utils.DrawText(screen, 0, y, "Press any key to continue...", styles.Default)
	y++
	utils.DrawText(screen, 0, y, "(q) quit", styles.Muted)
}
