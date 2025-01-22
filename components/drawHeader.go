package components

import (
	"main/styles"
	"main/utils"

	"github.com/gdamore/tcell/v2"
)

func DrawHeader(screen tcell.Screen, headerText string) {
	utils.DrawText(screen, 0, 0, " Monkey ", styles.Display)
	utils.DrawText(screen, 10, 0, headerText, styles.Default.Bold(true))
}