// Package minesweeper solves an Exercism challenge.
package minesweeper

import (
	"errors"
)

// Count updates the byte slice to make it a correct minesweeper board.
// Count returns an error if the input is malformed.
func (b Board) Count() error {
	// Every row must have the same length.
	l := len(b[0])
	for rowI, row := range b {
		if len(row) != l {
			return errors.New("rows are not of uniform length")
		}
		for colI, col := range row {
			// This first part is all validation.
			if rowI == 0 || rowI == len(b)-1 {
				if colI == 0 || colI == l-1 {
					if col != '+' {
						return errors.New("missing '+' in a corner")
					}
					continue
				}
				if col != '-' {
					return errors.New("missing '-' in top/bottom row")
				}
				continue
			}
			if colI == 0 || colI == l-1 {
				if col != '|' {
					return errors.New("missing '|' at edge of board")
				}
				continue
			}
			// Okay mostly done with validation.
			switch col {
			case '*':
				continue
			case ' ':
				num := 0
				// Don't have to worry about going out of bounds because of the
				// border, which we've already validated.
				for dr := -1; dr < 2; dr++ {
					for dc := -1; dc < 2; dc++ {
						if b[rowI+dr][colI+dc] == '*' {
							num++
						}
					}
				}
				if num != 0 {
					b[rowI][colI] = '0' + byte(num)
				}
			default:
				return errors.New("invalid character in board")
			}
		}
	}
	return nil
}
