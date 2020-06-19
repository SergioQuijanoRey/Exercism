// Package raindrops contains the function Convert which solves the given problem
package raindrops

import "strconv"

// Convert converts an integer to a string using the given rules
// If input has 3 as a factor, add 'Pling' to the result
// If input has 5 as a factor, add 'Plang' to the result
// If input has 7 as a factor, add 'Plong' to the result
// If input does not have any of 3, 5, or 7 as a factor, the result should be the digits of the number
func Convert(input int) string {
	result := ""

	if input%3 == 0 {
		result = result + "Pling"
	}

	if input%5 == 0 {
		result = result + "Plang"
	}

	if input%7 == 0 {
		result = result + "Plong"
	}

	if result == "" {
		result = strconv.Itoa(input)
	}

	return result
}
