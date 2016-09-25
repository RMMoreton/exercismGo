// Package anagram solves an Exercism challenge.
package anagram

import (
	"strings"
)

// Detect passes a subject and a slice of candidates, and
// returns a slice containing all candidates which are
// anagrams of subject.
func Detect(s string, c []string) []string {
	var res = make([]string, 0)

	s = strings.ToLower(s)
	sCount := count(s)

	for _, ci := range c {
		ci = strings.ToLower(ci)
		// if lengths aren't equal, it's 100% not an anagram
		// if the words are the same, it doesn't count as an anagram
		if len(s) != len(ci) || ci == s {
			continue
		}
		ciCount := count(ci)
		if equalCounts(sCount, ciCount) {
			res = append(res, ci)
		}
	}
	return res
}

// count passes a string and returns a map[rune]int of the
// rune counts in that string.
func count(s string) [26]int {
	var res [26]int
	for _, r := range s {
		res[r-'a']++
	}
	return res
}

// equalCounts returns true iff the two passed arrays are
// equal.
func equalCounts(s, c [26]int) bool {
	for i := 0; i < 26; i++ {
		if s[i] != c[i] {
			return false
		}
	}
	return true
}
