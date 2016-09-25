// Package trinary solves an Exercism challenge.
package trinary

import (
	"errors"
)

// ParseTrinary takes a string and parses it as a trinary number. On invalid
// input, the value returned is 0, and an error is returned as well.
func ParseTrinary(s string) (int64, error) {
	var result int64 = 0
	for _, c := range s {
		result *= 3
		switch c {
		case '0':
			continue
		case '1':
			result += 1
		case '2':
			result += 2
		default:
			return 0, errors.New("Invalid input found")
		}
		if result < 0 {
			return 0, errors.New("Overflow occured")
		}
	}
	return result, nil
}
