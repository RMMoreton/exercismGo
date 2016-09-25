// Package lsproduct solves an Exercism challenge.
package lsproduct

import (
	"errors"
	"fmt"
)

// testVersion is needed for the testing framework.
const testVersion = 3

// LargestSeriesProduct calculates the largest product of a series
// of length n in the passed number.
func LargestSeriesProduct(s string, n int) (int, error) {
	// sanity
	if n < 0 {
		return -1, errors.New("n may not be less than 1")
	}

	// turn s into a slice of ints
	var num = make([]int, 0, len(s))
	for _, r := range s {
		if r < '0' || r > '9' { // assume all digits in s are ASCII
			return -1, errors.New(fmt.Sprintf("string %s contains an unexpected rune", s))
		}
		num = append(num, int(r-'0'))
	}
	// more sanity
	if len(num) < n {
		return -1, errors.New(fmt.Sprintf("string %s does not contain %d digits", s, n))
	}

	// Calculate the product of the first n digits, using numZeroes to keep track of how
	// many zeroes are encountered. The product is only considered valid when numZeros
	// is equal to 0.
	curProd, maxProd, numZeroes := 1, 0, 0
	for i := 0; i < n; i++ {
		if num[i] == 0 {
			numZeroes++
		} else {
			curProd *= num[i]
		}
	}
	if numZeroes == 0 {
		maxProd = curProd
	}

	// Slide the 'product window' through num, incrementing and decrementing numZeroes
	// as appropriate to indicate when the product is valid.
	for i := n; i < len(num); i++ {
		if num[i] == 0 {
			numZeroes++
		} else {
			curProd *= num[i]
		}
		if num[i-n] == 0 {
			numZeroes--
		} else {
			curProd /= num[i-n]
		}
		if curProd > maxProd && numZeroes == 0 {
			maxProd = curProd
		}
	}

	return maxProd, nil
}
