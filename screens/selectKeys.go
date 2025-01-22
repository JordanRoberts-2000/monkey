package screens

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/inancgumus/screen"
)

func SelectKeys() {
	screen.Clear()
	screen.MoveTopLeft()

	display := color.New(color.FgWhite, color.BgGreen, color.Bold)
	boldText := color.New(color.Bold)
	display.Print(" Monkey ")

	boldText.Println("  How many keys to include?")
	fmt.Print(color.GreenString("\t  ⦿"))
	fmt.Println(" Letters only")
	color.RGB(63, 63, 70).Println("\t  ◯ Letters and spaces only")
	color.RGB(63, 63, 70).Println("\t  ◯ Letters, numbers, special characters: ,.!?")
	color.RGB(63, 63, 70).Println("\t  ◯ Every key on a macbook")

	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadByte()
}