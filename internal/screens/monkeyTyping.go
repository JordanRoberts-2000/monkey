package screens

import (
	"main/internal/styles"
	"main/internal/utils"

	"github.com/gdamore/tcell/v2"
)

func MonkeyTyping(screen tcell.Screen) {

	utils.DrawText(screen, 0, 2, "Total keys pressed: 100,000", styles.Default)
	utils.DrawText(screen, 0, 3, "\"the sentence i typed\"", styles.Muted)
	// color.RGB(63, 63, 70).Print("bieuldbewyafbsdzjkhvbufsdjakghfuslnjfksghusfhidnlvjbzkfguhsoilnkvjbzkdfguhesidlknjzbkvhfuhsijkvzbduhiflkncj`bdhuhilasfnbjksdbj`sndjbzsgjbufhdffbdsthisdfsufbhsydunjbdfhuhznjsbfhuzhnsjvbhuhsnjvbhushnjxvbhdushijvxbushcnjvxbcuhdjnxbudhsncdjvbkuhsjvxbkhudhscjvxbudihsjxvdfjcxvdbvbcxjhvbjdhbvjdnfhbjnvbduhxcnjbkhucjlnkvzhb")
	// reader := bufio.NewReader(os.Stdin)
}
