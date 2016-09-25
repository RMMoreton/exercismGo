// Package octal solves an Exercism challenge.
package octal

import (
	"errors"
)

// ParseOctal takes a string and returns the result of
// interpreting that string as an octal number.
func ParseOctal(s string) (int64, error) {
	var result int64
	for _, r := range s {
		if r < '0' || r > '7' {
			return 0, errors.New("encountered unexpected rune")
		}
		result *= 8
		result += int64(r - '0')
	}
	return result, nil
}
