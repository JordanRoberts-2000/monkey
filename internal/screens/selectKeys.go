package screens

import (
	"main/internal/styles"
	"main/internal/utils"

	"github.com/gdamore/tcell/v2"
)

func SelectKeys(screen tcell.Screen) {
	utils.DrawText(screen, 1, 2, "●Letters only", styles.Default)
	utils.DrawText(screen, 1, 3, "○Letters and spaces only", styles.Default)
	utils.DrawText(screen, 1, 4, "○Letters, numbers, special characters: ,.!?", styles.Default)
	utils.DrawText(screen, 1, 5, "○Every key on a macbook", styles.Default)
}
