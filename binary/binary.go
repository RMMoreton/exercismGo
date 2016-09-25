// Package binary solves an Exercism challenge.
package binary

import (
	"errors"
)

// ParseBinary passes a string representation of a
// binary number and returns the int it represents.
func ParseBinary(s string) (int, error) {
	if s == "" {
		return 0, errors.New("empty string is not parseable")
	}
	var res int
	for _, r := range s {
		res = res << 1
		switch r {
		case '1':
			res += 1
		case '0': // do nothing, we already did the shift
		default:
			return 0, errors.New("non-binary rune found in input string")
		}
	}
	return res, nil
}
