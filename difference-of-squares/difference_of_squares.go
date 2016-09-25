// Package diffsquares solves an Exercism challenge.
package diffsquares

// SquareOfSums returns the square of the sum of the first n natural numbers.
func SquareOfSums(n int) int {
	sum := (n*n + n) / 2
	return sum * sum
}

// SumOfSquares returns the sum of the squares of the first n natural numbers.
func SumOfSquares(n int) int {
	return (2*n*n*n + 3*n*n + n) / 6
}

// Difference returns SquareOfSums - SumOfSquares
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
