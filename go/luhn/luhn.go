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
	turn := false

	// We need to modify the input string, and strings in go are inmutable
	// So we mutate a *decoded* version of the string and return the *encoded
	// modified version
	decoded_input := []rune(input)

	for i := len(input) - 1; i >= 0; i-- {
		if turn {
			// We take the digit and double it
			digit, _ := strconv.Atoi(string(input[i]))
			digit = digit * 2
			if digit > 9 {
				digit = digit - 9
			}

			// We change the digit in the *decoded* string
			decoded_input[i] = rune(strconv.FormatInt(int64(digit), 10)[0])

		}

		turn = !turn
	}

	input = string(decoded_input)

	// Checking if the sum is divisible by 10
	sum := 0
	for _, char := range input {
		currentDigit, _ := strconv.Atoi(string(char))
		sum = sum + currentDigit
	}

	return sum%10 == 0

}
