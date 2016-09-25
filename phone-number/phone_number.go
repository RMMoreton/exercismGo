// Package phonenumber solves an Exercism challenge.
package phonenumber

import (
	"bytes"
	"errors"
	"unicode"
)

var errTooFew = errors.New("too few digits to be a phone number")
var errTooMany = errors.New("too many digits to be a phone number")

// Number assumes the input is some sort of phone number, and
// attempts to clean it up. It returns the number with no
// delimiters, or an error.
func Number(p string) (string, error) {
	var buff bytes.Buffer
	numWritten := 0
	for _, c := range p {
		if unicode.IsNumber(c) {
			buff.WriteRune(c)
			numWritten++
		}
	}
	s := buff.String()
	switch {
	case numWritten < 10:
		return "", errTooFew
	case numWritten == 10:
		return s, nil
	case numWritten == 11 && s[0] == byte('1'):
		return s[1:], nil
	case numWritten >= 11:
		return "", errTooMany
	default:
		panic("shouldn't be able to get here")
	}
}

// AreaCode takes a string and attempts to grab the area area
// code out of it (assuming it's a number)
func AreaCode(p string) (string, error) {
	p, err := Number(p)
	if err != nil {
		return "", err
	}
	return p[:3], nil
}

// Format takes a string and attempts to turn it in to a
// well formatted phone number.
func Format(p string) (string, error) {
	p, err := Number(p)
	if err != nil {
		return "", err
	}
	var buff bytes.Buffer
	buff.WriteRune('(')
	for _, c := range p[:3] {
		buff.WriteRune(c)
	}
	buff.WriteRune(')')
	buff.WriteRune(' ')
	for _, c := range p[3:6] {
		buff.WriteRune(c)
	}
	buff.WriteRune('-')
	for _, c := range p[6:] {
		buff.WriteRune(c)
	}
	return buff.String(), nil
}
