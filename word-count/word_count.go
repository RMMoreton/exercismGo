// Package wordcount solves an Exercism challenge.
package wordcount

import (
	"strings"
	"unicode"
)

// testVersion is required for the test suite.
const testVersion = 2

// Frequency will be used to map words to their word count
type Frequency map[string]int

// WordCount returns a map mapping each word in phrase
// to the number of times it appears.
func WordCount(phrase string) Frequency {
	var res = make(map[string]int)

	phrase = strings.ToLower(phrase)
	words := strings.FieldsFunc(phrase, normalize)

	for _, w := range words {
		if w == "" {
			continue
		}
		res[w]++
	}

	return res
}

// normalize returns true if the rune should be dropped from
// the normalized string, and false otherwise.
func normalize(r rune) bool {
	if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
		return true
	}
	return false
}
