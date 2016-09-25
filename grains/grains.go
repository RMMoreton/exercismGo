// Package grains solves an Exercism challenge.
package grains

import (
	"errors"
	"fmt"
)

// Square returns the number of grains of wheat on a given tile of a chess board.
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New(fmt.Sprintf("%d is not a valid square on a chessboard", n))
	}
	// The number of grains starts at 1 and then doubles, so the number of grains on
	// each square is a power of two. We can get powers of two very quickly with
	// bit shifting.
	return 1 << uint(n-1), nil
}

// Total returns the total number of grains of the chess board.
func Total() uint64 {
	var result uint64
	// 2^0 + 2^1 + ... + 2^n = 2^(n+1) - 1 which, written in binary, is n 1's. Since
	// I'm looking for the sum of the first 64 powers of two, I need the binary number
	// with 64 1's. I get that by taking the complement of a 64-bit 0.
	return ^result
}
