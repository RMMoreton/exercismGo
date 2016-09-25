// Package cryptosquare solves an Exercism challenge.
package cryptosquare

import (
	"bytes"
	"math"
	"strings"
	"unicode"
)

// testVersion is required by the testing framework
const testVersion = 2

// Encode encodes a string using the square cypher.
func Encode(p string) string {
	p = strings.ToLower(p)
	var buff bytes.Buffer
	for _, r := range p {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			buff.WriteRune(r)
		}
	}
	s := buff.String()

	c := int(math.Ceil(math.Sqrt(float64(len(s)))))

	buff = *new(bytes.Buffer)
	for cOffset := 0; cOffset < c; cOffset++ {
		for sPointer := cOffset; sPointer < len(s); sPointer += c {
			buff.WriteByte(s[sPointer]) // Everything should be ASCII by now...
		}
		if cOffset < c-1 {
			buff.WriteRune(' ')
		}
	}

	return buff.String()
}
