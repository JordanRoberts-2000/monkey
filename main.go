package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
)


func main() {
	// Initialize tcell screen
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("Failed to create tcell screen: %v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("Failed to initialize tcell screen: %v", err)
	}
	defer screen.Fini() 

	// loop
	running := true
	for running {
		screen.Clear()

		screen.SetContent(0,0,' ',nil, tcell.StyleDefault.Background(tcell.NewRGBColor(22, 163, 74)).Foreground(tcell.ColorWhiteSmoke.TrueColor()))
		screen.SetContent(1,0,'M',nil, tcell.StyleDefault.Background(tcell.NewRGBColor(50, 205, 50)).Foreground(tcell.ColorWhiteSmoke.TrueColor()))
		screen.SetContent(2,0,'o',nil, tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite))
		screen.SetContent(3,0,'n',nil, tcell.StyleDefault.Background(tcell.NewRGBColor(22, 163, 74)).Foreground(tcell.ColorWhiteSmoke.TrueColor()))
		screen.SetContent(4,0,'k',nil, tcell.StyleDefault.Background(tcell.NewRGBColor(22, 163, 74)).Foreground(tcell.ColorWhite))
		screen.SetContent(5,0,'e',nil, tcell.StyleDefault.Background(tcell.NewRGBColor(22, 163, 74)).Foreground(tcell.ColorWhite))
		screen.SetContent(6,0,'y',nil, tcell.StyleDefault.Background(tcell.NewRGBColor(22, 163, 74)).Foreground(tcell.ColorWhite))
		screen.SetContent(7,0,' ',nil, tcell.StyleDefault.Background(tcell.NewRGBColor(22, 163, 74)).Foreground(tcell.ColorWhite))

		screen.SetContent(10,0,'I',nil, tcell.StyleDefault.Attributes(tcell.AttrBold))
		screen.SetContent(11,0,'n',nil, tcell.StyleDefault.Attributes(tcell.AttrBold))
		screen.SetContent(12,0,'f',nil, tcell.StyleDefault.Attributes(tcell.AttrBold))
		screen.SetContent(13,0,'i',nil, tcell.StyleDefault.Attributes(tcell.AttrBold))
		screen.SetContent(14,0,'n',nil, tcell.StyleDefault.Attributes(tcell.AttrBold))
		screen.SetContent(15,0,'i',nil, tcell.StyleDefault.Attributes(tcell.AttrBold))
		screen.SetContent(16,0,'t',nil, tcell.StyleDefault.Attributes(tcell.AttrBold))
		screen.SetContent(17,0,'e',nil, tcell.StyleDefault.Attributes(tcell.AttrBold))
		screen.SetContent(18,0,' ',nil, tcell.StyleDefault.Attributes(tcell.AttrBold))
		screen.SetContent(19,0,'M',nil, tcell.StyleDefault.Attributes(tcell.AttrBold))
		screen.SetContent(20,0,'o',nil, tcell.StyleDefault.Attributes(tcell.AttrBold))
		screen.SetContent(21,0,'n',nil, tcell.StyleDefault.Attributes(tcell.AttrBold))
		screen.SetContent(22,0,'k',nil, tcell.StyleDefault.Attributes(tcell.AttrBold))
		screen.SetContent(23,0,'e',nil, tcell.StyleDefault.Attributes(tcell.AttrBold))
		screen.SetContent(24,0,'y',nil, tcell.StyleDefault.Attributes(tcell.AttrBold))
		screen.SetContent(25,0,' ',nil, tcell.StyleDefault.Attributes(tcell.AttrBold))
		screen.SetContent(26,0,'T',nil, tcell.StyleDefault.Attributes(tcell.AttrBold))
		screen.SetContent(27,0,'h',nil, tcell.StyleDefault.Attributes(tcell.AttrBold))
		screen.SetContent(28,0,'e',nil, tcell.StyleDefault.Attributes(tcell.AttrBold))
		screen.SetContent(29,0,'o',nil, tcell.StyleDefault.Attributes(tcell.AttrBold))
		screen.SetContent(30,0,'r',nil, tcell.StyleDefault.Attributes(tcell.AttrBold))
		screen.SetContent(31,0,'e',nil, tcell.StyleDefault.Attributes(tcell.AttrBold))
		screen.SetContent(32,0,'m',nil, tcell.StyleDefault.Attributes(tcell.AttrBold))

		screen.Show()
	
}
}
