package components

import (
	"main/internal/styles"
	"main/internal/utils"
	"strconv"

	"github.com/gdamore/tcell/v2"
)

func DrawHeader(screen tcell.Screen, headerText string, index int) {
	utils.DrawText(screen, 0, 0, " Monkey ", styles.Display)
	utils.DrawText(screen, 9, 0, "("+strconv.Itoa(index)+"/4) ", styles.Muted)

	utils.DrawText(screen, 15, 0, headerText, styles.Default.Bold(true))
}
