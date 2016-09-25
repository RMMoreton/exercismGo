// Package raindrops solves an Exercism challenge.
package raindrops

import (
	"fmt"
)

const testVersion = 2

// Convert passes an int and returns that int in "raindrop language."
func Convert(x int) string {
	var (
		rain bool = false
		str  string
	)
	if x%3 == 0 {
		str += fmt.Sprintf("Pling")
		rain = true
	}
	if x%5 == 0 {
		str += fmt.Sprintf("Plang")
		rain = true
	}
	if x%7 == 0 {
		str += fmt.Sprintf("Plong")
		rain = true
	}
	// The following section also solves the problem, but seems needlessly
	// complex to me. It was more interesting to code, though.
	/*
		type intStr struct{
			i int
			s string
		}
		m := make([]intStr, 3)
		m[0] = intStr{3, "Pling"}
		m[1] = intStr{5, "Plang"}
		m[2] = intStr{7, "Plong"}
		for _, st := range m {
			if x % st.i == 0 {
				str += st.s
				rain = true
			}
		}
	*/
	if !rain {
		str += fmt.Sprintf("%d", x)
	}
	return str
}
