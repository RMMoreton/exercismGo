// Package pythagorean solves an Exercism challenge.
package pythagorean

import (
	"math"
)

// Triplet holds a Pythagorean triplet
type Triplet [3]int

// Range returns all pythagorean tripples with all three components
// within min and max, inclusive.
func Range(min, max int) []Triplet {
	var res = make([]Triplet, 0)
	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			c := math.Floor(math.Sqrt(float64(a*a + b*b)))
			if int(c) <= max && c*c == float64(a*a+b*b) {
				res = append(res, Triplet{a, b, int(c)})
			}
		}
	}
	return res
}

// Sum returns all Pythagorean tripples whos components add to p.
func Sum(p int) []Triplet {
	var res = make([]Triplet, 0)
	pos := Range(1, p)
	for _, t := range pos {
		if t[0]+t[1]+t[2] == p {
			res = append(res, t)
		}
	}
	return res
}
