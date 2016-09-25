// Package beer solves an Exercism challenge.
package beer

import (
	"errors"
	"fmt"
	"strings"
)

// Verse returns a single verse of the song.
func Verse(n int) (string, error) {
	if n > 99 || n < 0 {
		return "", errors.New(fmt.Sprintf("invalid verse number %d", n))
	}
	switch {
	case n == 0:
		return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n", nil
	case n == 1:
		return "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n", nil
	case n == 2:
		return "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n", nil
	default:
		return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n", n, n, n-1), nil
	}
}

// Verses returns multiple verses of the song.
func Verses(upper, lower int) (string, error) {
	if upper > 99 || upper < 0 {
		return "", errors.New(fmt.Sprintf("invalid upper bound %d", upper))
	}
	if lower > 99 || lower < 0 {
		return "", errors.New(fmt.Sprintf("invalid lower bound %d", lower))
	}
	if upper < lower {
		return "", errors.New(fmt.Sprintf("invalid upper/lower pair %d, %d", upper, lower))
	}
	verses := make([]string, 0)
	for i := upper; i >= lower; i-- {
		next, _ := Verse(i)
		verses = append(verses, next)
	}
	toReturn := strings.Join(verses, "\n")
	toReturn += "\n"
	return toReturn, nil
}

// Song returns the whole song.
func Song() string {
	song, _ := Verses(99, 0)
	return song
}
