// Package romannumerals solves an Exercism challenge.
package romannumerals

import (
	"bytes"
	"errors"
	"fmt"
)

// testVersion is needed for testing.
const testVersion = 2

// Numeral is a structure to hold an integer and that integer's
// roman numeral equivalent.
type Numeral struct {
	arabic int
	roman  rune
}

// Conversion maps integers to their roman numeral counter-parts.
var Conversion = []Numeral{
	{1, 'I'},
	{5, 'V'},
	{10, 'X'},
	{50, 'L'},
	{100, 'C'},
	{500, 'D'},
	{1000, 'M'},
}

// ToRomanNumeral converts an integer < 4000 to a
// roman numeral.
func ToRomanNumeral(n int) (string, error) {
	// sanity
	if n <= 0 || n >= 4000 {
		return "", errors.New(fmt.Sprintf("converting %d is not possible", n))
	}

	var res bytes.Buffer
	for i := len(Conversion) - 1; i >= 2; i -= 2 {
		// Large is the current largest-possible numeral, medium
		// is one smaller, and small is two smaller.
		large := Conversion[i]
		medium := Conversion[i-1]
		small := Conversion[i-2]

		// Put a bunch of M's, C's, or X's on.
		for n >= large.arabic {
			res.WriteRune(large.roman)
			n -= large.arabic
		}

		switch {
		// Put on a CM, XC, or IX.
		case n-medium.arabic-4*small.arabic >= 0:
			res.WriteRune(small.roman)
			res.WriteRune(large.roman)
			n = n - medium.arabic - 4*small.arabic

		// Put on a D, L, or V.
		case n-medium.arabic >= 0:
			res.WriteRune(medium.roman)
			n = n - medium.arabic

		// Put on a CD, XL, or IV.
		case n-4*small.arabic >= 0:
			res.WriteRune(small.roman)
			res.WriteRune(medium.roman)
			n = n - 4*small.arabic
		}
	}

	// Put on as many I's as necessary.
	for n > 0 {
		res.WriteRune(Conversion[0].roman)
		n -= 1
	}

	return res.String(), nil
}
