// Package palindrome solves an Exercism challenge.
package palindrome

import (
	"errors"
	"strconv"
)

// Product stores a palindromic integer and all two-factor
// factorizations of it.
type Product struct {
	Product        int
	Factorizations [][2]int
}

// Products returns the minimum and the maximum Product's whose
// factors are within the given range.
func Products(min, max int) (Product, Product, error) {
	var resMin, resMax Product
	var found bool
	// sanity
	if min > max {
		return resMin, resMax, errors.New("fmin > fmax...")
	}
	// find all palindromic numbers with two-factor factorizations
	// in the specified range, keep track of the ones we want
	for i := min; i <= max; i++ {
		for j := i; j <= max; j++ {
			p := i * j
			if isPalindrome(p) {
				switch {
				case !found: // initialize because this is the first we've seen
					resMin = Product{p, [][2]int{{i, j}}}
					resMax = resMin
					found = true
				case p < resMin.Product: // new min
					resMin = Product{p, [][2]int{{i, j}}}
				case p > resMax.Product: // new max
					resMax = Product{p, [][2]int{{i, j}}}
				case p == resMin.Product: // new factorization of min
					resMin.Factorizations = append(resMin.Factorizations, [2]int{i, j})
				case p == resMax.Product: // new factorization of max
					resMax.Factorizations = append(resMax.Factorizations, [2]int{i, j})
				}
			}
		}
	}
	// must have initialized resMin and resMax
	if !found {
		return resMin, resMax, errors.New("No palindromes...")
	}

	return resMin, resMax, nil
}

// isPalindrome returns whether the passed in is
// palindromic (base 10).
func isPalindrome(n int) bool {
	var s = strconv.Itoa(n)
	var l = len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}
