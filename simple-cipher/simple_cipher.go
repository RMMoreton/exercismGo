// Package cipher solves an Exercism challenge.
package cipher

import (
	"strings"
)

// A Vigenere cipher is exactly what is sounds like.
type Vigenere struct {
	shift []byte
}

// NewCeaser returns a Ceaser Cipher.
func NewCaesar() Cipher {
	return &Vigenere{shift: []byte{3}}
}

// NewShift returns a shift cipher with the given shift value.
func NewShift(s int) Cipher {
	if s < -25 || s > 25 || s == 0 {
		return nil
	}
	return &Vigenere{shift: []byte{byte(s)}}
}

// NewVigenere returns a Vigenere cipher with the given string as it's key.
func NewVigenere(s string) Cipher {
	// Can't be made of all 'a's, and this will also return nil on the empty
	// string.
	if strings.Trim(s, "a") == "" {
		return nil
	}
	ret := Vigenere{shift: make([]byte, len(s))}
	for i, r := range []byte(s) {
		if r < 'a' || r > 'z' {
			return nil
		}
		ret.shift[i] = r - 'a'
	}
	return &ret
}

// Encoding means adding shift to each character of the input.
func (c *Vigenere) Encode(s string) string {
	s = strings.ToLower(s)
	ret := make([]byte, 0)
	i := 0
	for _, r := range []byte(s) {
		r -= 'a'
		// Invalid input returns empty string
		if r < 0 || r > 25 {
			continue
		}
		r += c.shift[i%len(c.shift)]
		// If it's less then 100, then just mod and we'll get within range.
		if r <= 100 {
			r %= 26
		} else {
			// We're greater then 100, so we wrapped around to around 255.
			// Adding 26 will wrap us back around to where we want to be.
			r += 26
		}
		r += 'a'
		ret = append(ret, r)
		i++
	}
	return string(ret)
}

// Decoding means subtracting shift from each character of the input.
func (c *Vigenere) Decode(s string) string {
	ret := make([]byte, 0)
	i := 0
	for _, r := range []byte(s) {
		r -= 'a'
		// Invalid input returns empty string
		if r < 0 || r > 25 {
			continue
		}
		r -= c.shift[i%len(c.shift)]
		// Similar to the encode function.
		if r <= 100 {
			r %= 26
		} else {
			r += 26
		}
		r += 'a'
		ret = append(ret, r)
		i++
	}
	return string(ret)
}
