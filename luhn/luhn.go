// Package luhn solves an Exercism challenge.
package luhn

import (
	"bytes"
	"fmt"
	"strconv"
	"unicode"
)

// Valid returns true if the passed string is valid
// according to the luhn checksum formula.
func Valid(s string) bool {
	n := convert(s)
	if n < 0 {
		return false
	}
	return luhnRemainder(n, 0) == 0
}

// AddCheck takes a raw number and appends a digit so that
// the number is a valid luhn checksum.
func AddCheck(s string) string {
	n := convert(s) // first convert the original to make sure there's a number somewhere
	if n < 0 {
		return ""
	}
	r := luhnRemainder(n, 1)
	return s + strconv.Itoa((10-r)%10) // mod by 10 in case r == 0
}

// convert passes a string and converts that string
// into an unsigned integer, ignoring all non-integer runes.
// A negative return value indicates the string contains no digits.
func convert(s string) int {
	var buff bytes.Buffer
	for _, r := range s {
		if unicode.IsNumber(r) {
			buff.WriteRune(r)
		}
	}
	s = buff.String()
	if s == "" {
		return -1
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("Atoi(%s) failed", s))
	}
	return n
}

// luhnRemainder calculates the remainder of running the Luhn
// checksum algorithm on the passed integer
func luhnRemainder(n, double int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		n /= 10
		if double == 1 {
			digit *= 2
			if digit >= 10 {
				digit -= 9
			}
		}
		sum += digit
		double = (double + 1) % 2
	}
	return sum % 10
}
