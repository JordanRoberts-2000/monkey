package utils

import (
	"main/styles"

	"github.com/gdamore/tcell/v2"
)

func DrawHeader(screen tcell.Screen, headerText string) {
	DrawText(screen, 0, 0, " Monkey ", styles.Display)
	DrawText(screen, 10, 0, headerText, styles.Default.Bold(true))
}