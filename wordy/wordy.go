// Package wordy solves an Exercism challenge.
package wordy

import (
	"strings"
)

// Operator constants used to indicate what type of
// operation to perform.
const (
	ADD  = iota
	SUB  = iota
	MULT = iota
	DIV  = iota
)

// opPhrasePair holds an operation and the corresponding
// phrase.
type opPhrasePair struct {
	op     int
	phrase string
}

// opPhrasePairs holds all pairs of operations and phrases.
// I was using a map for this, but the map slowed things
// down, and I don't need all the fancy map operations.
// This is fast, and keeps the operation paired with it's
// phrase.
var opPhrasePairs = []opPhrasePair{
	{ADD, "plus"},
	{SUB, "minus"},
	{MULT, "multiplied by"},
	{DIV, "divided by"},
}

// Answer takes a word problem of the form "What is [x]
// [op] [y] ...?" and returns the result and whether
// the result is valid (i.e. the returned bool will be
// false on malformed input).
func Answer(q string) (int, bool) {
	q, ok := stripPrefix(q)
	if !ok {
		return 0, false
	}
	q, acc, ok := parseInt(q)
	if !ok {
		return 0, false
	}
	for q != "?" {
		var op, opperand int
		q, op = getOp(q)
		if op == -1 {
			return 0, false
		}
		q, opperand, ok = parseInt(q)
		if !ok {
			return 0, false
		}
		switch op {
		case ADD:
			acc += opperand
		case SUB:
			acc -= opperand
		case MULT:
			acc *= opperand
		case DIV:
			acc /= opperand
		default:
			return 0, false
		}
	}
	return acc, true
}

// stripPrefix takes a string s which should start with
// optional whitespace followed by
// "What is ", and returns that string with the whitespace
// and prefix
// taken off. If s does not have the correct prefix,
// the returned bool will be false.
func stripPrefix(s string) (string, bool) {
	s = strings.TrimSpace(s)
	res := strings.TrimPrefix(s, "What is")
	if res == s {
		return "", false
	}
	return res, true
}

// parseInt takes a string which should start with optianl
// whitespace followed by an
// integer, and returns the string with that whitespace
// and integer
// removed, that integer, and a bool to indicate whether
// or not the string started with an integer. In the event
// that s does not start with an integer, s is returned
// unchanged, the int is 0, and the bool is false.
func parseInt(s string) (string, int, bool) {
	s = strings.TrimSpace(s)
	index := 0
	num := 0
	neg := false
	if s[0] == '-' {
		neg = true
		index++
	}
	for ; s[index] >= '0' && s[index] <= '9'; index++ {
		num *= 10
		num += int(s[index] - '0')
	}
	if neg {
		num *= -1
	}
	// Did s start with a number?
	if index == 0 || (index == 1 && neg) {
		return s, 0, false
	}
	return s[index:], num, true
}

// getOp takes a string, strips whitespace, and looks
// for an operation phrase (e.g. "multiplied by"), strips
// that phrase, and returns the string and the operation.
// If no valid op-phrase is found, the returned operation
// is -1.
func getOp(s string) (string, int) {
	s = strings.TrimSpace(s)
	for _, pair := range opPhrasePairs {
		if strings.HasPrefix(s, pair.phrase) {
			s = strings.TrimPrefix(s, pair.phrase)
			return s, pair.op
		}
	}
	return s, -1
}
