// Package hamming solves one of Exercism's Go challenges.
package hamming

import (
	"fmt"
)

// A simple constant declaration.
const testVersion = 4

// Error is used to indicate that two strings have different length.
type Error struct {
	a, b string
}

// Error's use is obvious.
func (e Error) Error() string {
	return fmt.Sprintf("strings %s and %s have different length", e.a, e.b)
}

// Distance computes the hamming distance between two DNA strands.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, Error{a, b}
	}
	var dist int
	for i := range a {
		if a[i] != b[i] {
			dist += 1
		}
	}
	return dist, nil
}
