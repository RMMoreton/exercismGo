// Package bob is a solution to an Exercism challenge.
package bob

import (
	"strings"
)

const testVersion = 2 // same as targetTestVersion

// Hey passes a string and returns Bob's response.
func Hey(q string) string {
	q = strings.TrimSpace(q)
	switch {
	case q == "":
		return "Fine. Be that way!"
	case strings.ToUpper(q) == q && strings.ToLower(q) != q:
		return "Whoa, chill out!"
	case q[len(q)-1] == '?':
		return "Sure."
	default:
		return "Whatever."
	}
}
