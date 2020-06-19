// Package collatzconjecture has the function CollatzConjecture which solves the problem
package collatzconjecture

import "fmt"

// CollCollatzConjecture calculates the number of steps needed to go from
// val to 1, using the procedure of the Collatz Conjecture
func CollatzConjecture(val int) (int, error) {
	if val <= 0 {
		return 0, fmt.Errorf("val has to be a positive integer")
	}

	steps := 0
	for val != 1 {
		if val%2 == 0 {
			val = val / 2
		} else {
			val = val*3 + 1
		}

		steps++
	}

	return steps, nil
}
