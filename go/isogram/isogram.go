// Package isogram has the function IsIsogram which tests if a given word is an isogram
package isogram

import "unicode"

// IsIsogram checks if a given word is an isogram
func IsIsogram(word string) bool {
	// For having a record of all visited letters of the word
	prevLetters := make(map[rune]bool)

	for _, letter := range word {
		if letter == ' ' || letter == '-' {
			continue
		}

		letter = unicode.ToLower(letter)
		if prevLetters[letter] {
			return false
		}

		prevLetters[letter] = true
	}

	return true
}
