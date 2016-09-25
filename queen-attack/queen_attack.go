// Package queenattack solves an Exercism challenge.
package queenattack

import (
	"errors"
)

// CanQueenAttack passes the position of two queens, and returns whether
// those queens can attack each other.
func CanQueenAttack(w, b string) (bool, error) {
	// First do all the verifying and converting of the passed strings.
	if len(w) != 2 || len(b) != 2 {
		return false, errors.New("both arguments to CanQueenAttack() must be of length 2")
	}
	if w == b {
		return false, errors.New("two pieces may not occupy the same space")
	}
	if !constrainToBoard(w) || !constrainToBoard(b) {
		return false, errors.New("all pieces must be on the board")
	}

	// Now decide whether the queens can attack.
	wL, wN := w[0], w[1]
	bL, bN := b[0], b[1]
	switch {
	case wL == bL:
		fallthrough
	case wN == bN:
		fallthrough
	case wL-bL == wN-bN || wL-bL == bN-wN:
		return true, nil
	}
	return false, nil
}

// constrainToBoard ensures that a two-byte string represents a valid position
// on a chess board.
func constrainToBoard(s string) bool {
	letter, num := s[0], s[1]
	if letter < 'a' || letter > 'h' {
		return false
	}
	if num < '1' || num > '8' {
		return false
	}
	return true
}
