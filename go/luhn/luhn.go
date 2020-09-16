// Package luhn has the function Valid which solves the given problem
package luhn

import (
	"strconv"
	"strings"
)

// Valid takes an string and returns a boolean representating wether is valid or
// invalid regarding luhn algorithm
func Valid(input string) bool {
	// Stripping white spaces
	input = strings.ReplaceAll(input, " ", "")

	// Strings of length 1 or less are not valid
	if len(input) <= 1 {
		return false
	}

	// Doubling second digits starting from the right
	// turn controls if a digit has to be doubled or not
	turn := len(input)%2 == 0

	// Calculating the sum of the string as specified in the problem
	sum := 0
	for _, char := range input {
		// Getting the digit and checking for errors
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			return false
		}

		// Doubling the digit
		if turn {
			// We take the digit and double it
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		// Calculating the sum
		sum += digit

		// Changing the turn
		turn = !turn
	}

	return sum%10 == 0

}
