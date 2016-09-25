// Package sieve solves an Exercism challenge.
package sieve

// Sieve passes an integer n and returns all primes
// less than or equal to n.
func Sieve(n int) []int {
	composite := make([]bool, n+1, n+1)
	primes := make([]int, 0, 0)
	for i := 2; i <= n; i++ {
		if !composite[i] {
			primes = append(primes, i)
			for j := i + i; j <= n; j += i {
				composite[j] = true
			}
		}
	}
	return primes
}
