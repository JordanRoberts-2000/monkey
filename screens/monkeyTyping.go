package screens

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/inancgumus/screen"
)

func MonkeyTyping() {
	screen.Clear()
	screen.MoveTopLeft()

	display := color.New(color.FgWhite, color.BgGreen, color.Bold)
	boldText := color.New(color.Bold)
	display.Print(" Monkey ")
	
	boldText.Println("  Monkey is mashing keys...")
	fmt.Println("\nTotal keys pressed: 100,000")
	color.RGB(63, 63, 70).Println("\"the sentence i typed\"")
	println()
	color.RGB(63, 63, 70).Print("bieuldbewyafbsdzjkhvbufsdjakghfuslnjfksghusfhidnlvjbzkfguhsoilnkvjbzkdfguhesidlknjzbkvhfuhsijkvzbduhiflkncj`bdhuhilasfnbjksdbj`sndjbzsgjbufhdffbdsthisdfsufbhsydunjbdfhuhznjsbfhuzhnsjvbhuhsnjvbhushnjxvbhdushijvxbushcnjvxbcuhdjnxbudhsncdjvbkuhsjvxbkhudhscjvxbudihsjxvdfjcxvdbvbcxjhvbjdhbvjdnfhbjnvbduhxcnjbkhucjlnkvzhb")
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadByte()
}