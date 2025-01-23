package screens

import (
	"main/internal/styles"
	"main/internal/utils"

	"github.com/gdamore/tcell/v2"
)

func EnterInput(screen tcell.Screen, userInput string, errorMessage string, optionIndex int) {
	charSet := ""
	switch optionIndex {
	case 0:
		charSet = "[Letters only]"
	case 1:
		charSet = "[Letters and spaces only]"
	case 2:
		charSet = "[Letters, numbers, special characters: ,.!?]"
	case 3:
		charSet = "[Every key on a MacBook]"
	}

	utils.DrawText(screen, 0, 3, "Target Input: ", styles.Default)
	utils.DrawText(screen, 14, 3, userInput, styles.Highlighted)
	utils.DrawText(screen, 0, 2, charSet, styles.Muted)
	utils.DrawText(screen, 0, 4, errorMessage, styles.Invalid)
}
