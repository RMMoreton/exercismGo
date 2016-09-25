// Package triangle is a solution to an Exercism challenge.
package triangle

import (
	"math"
)

type Kind int

// Const declarations for use by the tester.
const (
	testVersion = 2

	NaT Kind = iota
	Equ Kind = iota
	Iso Kind = iota
	Sca Kind = iota
)

// KindFromSides passes three triangle sides and returns the kind of
// triangle that these sides represent. No side may be 0.
func KindFromSides(a, b, c float64) Kind {
	var max float64

	// a, b, c must all be positive real numbers
	for _, val := range [3]float64{a, b, c} {
		if math.IsNaN(val) {
			return NaT
		}
		if math.IsInf(val, 0) {
			return NaT
		}
		if val <= 0 {
			return NaT
		}
		// Use this opportunity to grab the max side
		if val > max {
			max = val
		}
	}
	// a + b + c - max must be at least max by the triangle inequality
	if a+b+c-max < max {
		return NaT
	}

	// All equal means equilateral
	if a == b && b == c {
		return Equ
	}

	// Two equal means isosceles
	if a == b || a == c || b == c {
		return Iso
	}

	return Sca
}
