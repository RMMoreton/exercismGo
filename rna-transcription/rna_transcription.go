// Package strand solves an Exercism challenge.
package strand

import (
	"fmt"
	"strings"
)

// testVersion is required for testing.
const testVersion = 3

// ToRNA returns the RNA equivilent of a given DNA
// strand, which is expected to be correctly formatted.
func ToRNA(s string) string {
	return strings.Map(rnaMapper, s)
}

// rnaMapper returns the RNA equivilent of a given
// DNA molecule.
func rnaMapper(r rune) rune {
	switch r {
	case 'G':
		return 'C'
	case 'C':
		return 'G'
	case 'T':
		return 'A'
	case 'A':
		return 'U'
	default:
		panic(fmt.Sprintf("rune %q is not one of G, C, T, A", r))
	}
}
