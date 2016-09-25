// Package igpay solves an Exercism challenge.
package igpay

import (
	"bytes"
	"strings"
)

var vowels = map[byte]bool{
	'a': true, 'e': true, 'i': true,
	'o': true, 'u': true,
}

// PigLatin takes a string and turns every word in that
// string into pig latin. The string is assumed to be
// plain english text, lower case.
func PigLatin(s string) string {
	var buff bytes.Buffer
	words := strings.Split(s, " ")
	for i, w := range words {
		if i != 0 {
			buff.WriteRune(' ')
		}
		switch {
		case w[0] == 'y' || w[0] == 'x':
			if IsVowel(w[1]) {
				buff.WriteString(w[1:])
				buff.WriteByte(w[0])
			} else {
				buff.WriteString(w)
			}
		case IsVowel(w[0]):
			buff.WriteString(w)
		case w[0] == 'q' && w[1] == 'u':
			buff.WriteString(w[2:])
			buff.WriteString(w[:2])
		default:
			i := 1
			for !IsVowel(w[i]) {
				i++
				if w[i] == 'u' && w[i-1] == 'q' {
					i++
				}
			}
			buff.WriteString(w[i:])
			buff.WriteString(w[:i])
		}
		buff.WriteString("ay")
	}
	return buff.String()
}

// IsVowel takes a string (single english letter) and returns
// whether that string is a vowel.
func IsVowel(s byte) bool {
	if _, ok := vowels[s]; ok {
		return true
	}
	return false
}
