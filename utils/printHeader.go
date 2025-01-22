package utils

import (
	"fmt"
	"main/styles"
)

func PrintHeader(headerText string) {
	styles.Display.Print(" Monkey ")
	styles.BoldText.Println("  "+headerText)
	fmt.Println()
}