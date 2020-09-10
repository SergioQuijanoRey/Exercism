// Package diffsquares has functions to calculate different squares of series
package diffsquares

// SquareOfSum calculates the square of the natural series up to n
func SquareOfSum(n int) int {
	sum := (n * (n + 1)) / 2
	return sum * sum
}

// SquareOfSum calculates the natural squares series up to n
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

// Difference calculates the difference of the square of natural series up to n
// minus the natural squares series up to n
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
