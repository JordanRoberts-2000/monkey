package screens

import (
	"main/internal/styles"
	"main/internal/utils"

	"github.com/gdamore/tcell/v2"
)

func SelectKeys(screen tcell.Screen) {
	utils.DrawText(screen, 5, 5, "test", styles.Default)
	// utils.ClearScreen()
	// utils.PrintHeader("How many keys to include?")

	// styles.Highlighted.Print(" ●")
	// fmt.Println(" Letters only")
	// color.RGB(63, 63, 70).Println(" ○ Letters and spaces only")
	// color.RGB(63, 63, 70).Println(" ○ Letters, numbers, special characters: ,.!?")
	// color.RGB(63, 63, 70).Println(" ○ Every key on a macbook")

	// reader := bufio.NewReader(os.Stdin)
	// _, _ = reader.ReadByte()
}
