package utils

import (
	"main/internal/state"
	"strings"
)

func ValidateInput(appState *state.AppState) bool {
	allowedChars, exists := state.CharSets[appState.CharSetIndex]
	if !exists {
		return false
	}

	for _, char := range appState.UserInput {
		if !strings.ContainsRune(allowedChars, char) {
			if char == ' ' {
				appState.InvalidMessage = "invalid: spaces not allowed"
			} else {
				appState.InvalidMessage = "invalid character: " + string(char)
			}
			return false
		}
	}

	return true
}
