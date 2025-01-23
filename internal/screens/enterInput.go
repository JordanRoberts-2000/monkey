package screens

import (
	"main/internal/styles"
	"main/internal/utils"

	"github.com/gdamore/tcell/v2"
)

func EnterInput(screen tcell.Screen, userInput string) {
	utils.DrawText(screen, 0, 2, "Target Input: ", styles.Default)
	utils.DrawText(screen, 14, 2, userInput, styles.Highlighted)

	utils.DrawText(screen, 0, 3, "(B) back", styles.Muted)
}
