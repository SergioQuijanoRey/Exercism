// Package isogram has the function IsIsogram which tests if a given word is an isogram
package isogram

import "strings"

// IsIsogram checks if a given word is an isogram
func IsIsogram(word string) bool {
	// Normalize the given word
	word = strings.ToLower(word)
	word = strings.ReplaceAll(word, " ", "")
	word = strings.ReplaceAll(word, "-", "")

	// For having a record of all visited letters of the word
	prev_letters := make(map[rune]bool)

	for _, letter := range word {
		if _, found := prev_letters[letter]; found == true {
			return false
		}

		prev_letters[letter] = true
	}

	return true
}
