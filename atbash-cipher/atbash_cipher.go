// Package atbash solves an Exercism challenge.
package atbash

import (
	"bytes"
	"strings"
	"unicode"
)

// Atbash returns string s encrypted using the Atbash
// cypher.
func Atbash(s string) string {
	var buff bytes.Buffer
	s = strings.ToLower(s)
	numWritten := 0
	for _, c := range s {
		if (c >= 'a' && c <= 'z') || unicode.IsNumber(c) {
			if numWritten%5 == 0 && numWritten != 0 {
				buff.WriteRune(' ')
			}
			buff.WriteRune(Opposite(c))
			numWritten++
		}
	}
	return buff.String()
}

// Opposite returns the opposite (e.g. 'a' -> 'z') of rune r if
// r is a lower case lattin letter. If it is not, Opposite
// returns r.
func Opposite(r rune) rune {
	// If r has no opposite, return r.
	if r < 'a' || r > 'z' {
		return r
	}
	val := int(r - 'a')
	val = 25 - val
	return rune(val + int('a'))
}
