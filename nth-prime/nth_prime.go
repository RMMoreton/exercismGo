// Package prime solves an Exercism challenge.
package prime

// Nth returns the nth prime, or an error.
func Nth(n int) (int, bool) {
	if n < 1 {
		return 0, false
	}
	primes := make([]int, n)
	primes[0] = 2
	numFound := 1
	testPrimality := 3
	for numFound < n {
		for _, p := range primes {
			if p*p > testPrimality {
				primes[numFound] = testPrimality
				numFound++
				break
			}
			if testPrimality%p == 0 {
				break
			}
		}
		testPrimality++
	}
	return primes[n-1], true
}
