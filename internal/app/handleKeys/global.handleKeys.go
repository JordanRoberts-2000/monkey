package handlekeys

import (
	"main/internal/state"

	"github.com/gdamore/tcell/v2"
)

func GlobalKeys(state *state.AppState, ev *tcell.EventKey) {
	switch ev.Key() {
	case tcell.KeyEscape, tcell.KeyCtrlC, tcell.KeyCtrlQ, tcell.KeyCtrlX, tcell.KeyCtrlD:
		state.Running = false
		return
	}

	switch ev.Rune() {
	case 'q':
		state.Running = false
		return
	}
}
