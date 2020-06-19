// Package hamming contains the function Distance which calculates the hamming
// distance of two strings
package hamming

import (
	"fmt"
)

// Distance calculates the Hamming distance between two strings
// The two strings must have same lenght
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("Expecting two strings with same lenght, got strings of lenght %d and %d", len(a), len(b))
	}

	// Hamming distance is calculated
	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
