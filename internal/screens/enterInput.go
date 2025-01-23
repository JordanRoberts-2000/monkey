package screens

import (
	"main/internal/styles"
	"main/internal/utils"

	"github.com/gdamore/tcell/v2"
)

func EnterInput(screen tcell.Screen) {

	utils.DrawText(screen, 0, 2, "Target Input: ", styles.Default)
	utils.DrawText(screen, 0, 3, "(B) back", styles.Muted)
}
