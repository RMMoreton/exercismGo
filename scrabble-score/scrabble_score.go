// Package scrabble solves an Exercism challenge.
package scrabble

import (
	"strings"
)

const testVersion = 4

// scoreMap maps lowercase letters to integers based on
// the Scrabble score for those letters.
var scoreMap = map[rune]int{
	'a': 1, 'b': 3, 'c': 3, 'd': 2, 'e': 1,
	'f': 4, 'g': 2, 'h': 4, 'i': 1, 'j': 8,
	'k': 5, 'l': 1, 'm': 3, 'n': 1, 'o': 1,
	'p': 3, 'q': 10, 'r': 1, 's': 1, 't': 1,
	'u': 1, 'v': 4, 'w': 4, 'x': 8, 'y': 4,
	'z': 10,
}

// Score returns the Scrabble score of the passed word.
func Score(in string) int {
	in = strings.TrimSpace(in)
	var sum = 0
	if len(in) == 0 {
		return sum
	}
	for _, c := range strings.ToLower(in) {
		sum += scoreMap[c]
	}
	return sum
}
