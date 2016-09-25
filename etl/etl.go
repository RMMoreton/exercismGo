// Package etl solves an Exercism challenge.
package etl

import (
	"strings"
)

// Transform passes data in the old format and returns
// data in the new format.
func Transform(in map[int][]string) map[string]int {
	var out = make(map[string]int)
	for val, slice := range in {
		for _, letter := range slice {
			out[strings.ToLower(letter)] = val
		}
	}
	return out
}
