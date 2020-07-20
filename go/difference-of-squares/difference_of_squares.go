// Package diffsquares has functions to calculate different squares of series
package diffsquares

// SquareOfSum calculates the square of the natural series up to n
func SquareOfSum(n int) int {
	sum := 0.5 * float64(n) * float64(n+1)
	return int(sum * sum)
}

// SquareOfSum calculates the natural squares series up to n
func SumOfSquares(n int) int {
	// The factor is an aproximation of the value 1 / 6
	var factor float64 = 0.1666666666667
	return int(factor * float64(n) * float64(n+1) * float64(2*n+1))
}

// DDifference calculates the difference of the square of natural series up to n
// minus the natural squares series up to n
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
