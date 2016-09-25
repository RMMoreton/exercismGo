// Package dna solves an Exercism challenge.
package dna

import (
	"errors"
)

// Histogram holds the counts of each nucleotide
// in a DNA string.
type Histogram map[byte]int

// dnaHolder holds a histogram of counts, and the
// DNA string that was used to create that histogram.
type dnaHolder struct {
	s string
	h Histogram
}

// DNA turns a string into a dnaHolder. DNA does not
// complain if given a string with invalid nucleotides,
// it just ignores them.
func DNA(s string) dnaHolder {
	hist := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for i := range s {
		if !validNucleotide(s[i]) {
			continue
		}
		hist[s[i]] += 1
	}
	return dnaHolder{s, hist}
}

// Count returns the number of nucleotides of the
// passed type are in the given histogram.
func (d dnaHolder) Count(b byte) (int, error) {
	if !validNucleotide(b) {
		return 0, errors.New("invalid nucleotide passed to Count()")
	}
	return d.h[b], nil
}

// Counts returns the histogram associated with the
// given dnaHolder.
func (d dnaHolder) Counts() Histogram {
	return d.h
}

// validNucleotide returns true if the passed byte is
// a valid nucleotide, false otherwise.
func validNucleotide(b byte) bool {
	if b == 'A' || b == 'C' || b == 'G' || b == 'T' {
		return true
	}
	return false
}
