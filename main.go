package main

import (
	"errors"
	"log"
	"main/screens"
	"os"
	"unicode"

	"golang.org/x/term"
)

func validateInput(input string) error {
	for _, char := range input {
		if !unicode.IsLetter(char) || unicode.IsUpper(char) {
			return errors.New("string is invalid")
		}
	}
	return nil
}


func main() {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		log.Fatalf("Failed to get terminal size: %v", err)
	}

	screens.IntroScreen(height, width)
	screens.SelectKeys()
	screens.MonkeyTyping()
	// 	err := validateInput(input)
	// 	if err != nil {
	// 		fmt.Println("Error:", err)
	// 		fmt.Println("Please try again.")
	// 	} else {
	// 		break
	// 	}
	// }

	// seed := int64(42) // Your desired seed value
  // rng := rand.New(rand.NewSource(seed))
	// randomLetter := rune(rng.Intn(26) + 'a')
	// fmt.Printf("Random lowercase letter: %c\n", randomLetter)
}
