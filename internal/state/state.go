package state

type ScreenId int
type CharSetId int

const (
	Intro ScreenId = iota
	SelectKeys
	MonkeyTyping
)

var CharSets = map[CharSetId]string{
	0: "abcdefghijklmnopqrstuvwxyz",
	1: "abcdefghijklmnopqrstuvwxyz ",
	2: "abcdefghijklmnopqrstuvwxyz0123456789,!.?",
	3: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()_+-=[]{}|;:,.<>?/`~'\"",
}

type AppState struct {
	CurrentScreen ScreenId
	Running       bool
	CharSetIndex  CharSetId
}

func Initialize() *AppState {
	return &AppState{
		CurrentScreen: Intro,
		Running:       true,
		CharSetIndex:  0,
	}
}
