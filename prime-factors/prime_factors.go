// Package prime solves an Exercism challenge.
package prime

// testVersion is required for testing.
const testVersion = 2

// Factors returns a slice of all the factors of n.
func Factors(n int64) []int64 {
	toReturn := make([]int64, 0)
	var i int64
	for i = 2; n != 1; i++ {
		for n%i == 0 {
			n /= i
			toReturn = append(toReturn, i)
		}
	}
	return toReturn
}
