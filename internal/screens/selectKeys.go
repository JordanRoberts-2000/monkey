package screens

import (
	"main/internal/styles"
	"main/internal/utils"

	"github.com/gdamore/tcell/v2"
)

func SelectKeys(screen tcell.Screen, optionIndex int) {
	options := []string{
		"Letters only",
		"Letters and spaces only",
		"Letters, numbers, special characters: ,.!?",
		"Every key on a MacBook",
	}

	for i, option := range options {
		if i == optionIndex {
			utils.DrawText(screen, 1, 2+i, "●", styles.Highlighted)
			utils.DrawText(screen, 4, 2+i, option, styles.Default)
		} else {
			utils.DrawText(screen, 1, 2+i, "○"+option, styles.Muted)
		}
	}
}
