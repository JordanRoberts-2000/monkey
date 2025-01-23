package utils

import (
	"main/internal/state"
	"strings"
	"unicode"
)

func ValidateInput(appState *state.AppState) bool {
	allowedChars, exists := state.CharSets[appState.CharSetIndex]
	if !exists {
		appState.InvalidMessage = "invalid: invalid CharSetIndex"
		return false
	}

	hasValidChar := false
	for _, char := range appState.UserInput {
		if !strings.ContainsRune(allowedChars, char) {
			if char == ' ' {
				appState.InvalidMessage = "invalid: spaces not allowed"
			} else {
				appState.InvalidMessage = "invalid character: " + string(char)
			}
			return false
		}

		if !unicode.IsSpace(char) {
			hasValidChar = true
		}
	}

	if !hasValidChar {
		appState.InvalidMessage = "required to continue"
		return false
	}

	appState.InvalidMessage = ""
	return true
}
