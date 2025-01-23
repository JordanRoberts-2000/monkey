package styles

import "github.com/gdamore/tcell/v2"

var (
	Default = tcell.StyleDefault.Foreground(tcell.ColorWhite)
	Display = tcell.StyleDefault.
		Foreground(tcell.ColorBlack).
		Background(tcell.ColorLightGreen).
		Bold(true)
	Highlighted = tcell.StyleDefault.Foreground(tcell.ColorLightGreen)
	Invalid     = tcell.StyleDefault.Foreground(tcell.ColorIndianRed)
	Muted       = tcell.StyleDefault.Foreground(tcell.NewRGBColor(163, 163, 163))
)
