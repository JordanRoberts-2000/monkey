package state

type ScreenId int

const (
	Intro ScreenId = iota
	SelectKeys
	MonkeyTyping
)

type AppState struct {
	CurrentScreen    ScreenId
	Running  bool
}

func Initialize() (*AppState) {
	return &AppState{
		CurrentScreen: Intro,
		Running:       true,
	}
}