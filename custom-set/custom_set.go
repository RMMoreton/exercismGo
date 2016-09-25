// Package stringset solves an Exercism challenge.
package stringset

import (
	"bytes"
	"fmt"
)

// testVersion is needed for testing.
const testVersion = 3

// A Set holds a map of elements to booleans. The
// booleans don't matter, we're just going to test
// whether an element is in the maps key set.
type Set map[string]bool

// New returns a new Set.
func New() Set {
	return make(map[string]bool)
}

// NewFromSlice returns a new Set that has every
// element of slice sl in it.
func NewFromSlice(sl []string) Set {
	toReturn := New()
	for _, s := range sl {
		toReturn.Add(s)
	}
	return toReturn
}

// Add adds string st to set s.
func (s Set) Add(st string) {
	if _, ok := s[st]; !ok {
		s[st] = true
	}
}

// Delete removes string st from set s.
func (s Set) Delete(st string) {
	if _, ok := s[st]; ok {
		delete(s, st)
	}
}

// Has returns true iff string st is in set s.
func (s Set) Has(st string) bool {
	_, ok := s[st]
	return ok
}

// IsEmpty returns true iff set s is empty.
func (s Set) IsEmpty() bool {
	return s.Len() == 0
}

// Len returns how many elements are in set s.
func (s Set) Len() int {
	return len(s)
}

// Slice returns the elements of s in a slice.
func (s Set) Slice() []string {
	toReturn := make([]string, s.Len())
	i := 0
	for st := range s {
		toReturn[i] = st
		i++
	}
	return toReturn
}

// String returns the elements of s in a single string.
func (s Set) String() string {
	slice := s.Slice()
	var buff bytes.Buffer
	buff.WriteRune('{')
	for i, s := range slice {
		buff.WriteString(fmt.Sprintf("%q", s))
		if i < len(slice)-1 {
			buff.WriteString(", ")
		}
	}
	buff.WriteRune('}')
	return buff.String()
}

// Equal returns true iff Set s1 equals Set s2.
func Equal(s1, s2 Set) bool {
	if s1.Len() != s2.Len() {
		return false
	}
	for k := range s1 {
		if _, ok := s2[k]; !ok {
			return false
		}
	}
	return true
}

// Subset returns true iff Set s1 is a subset of Set s2.
func Subset(s1, s2 Set) bool {
	if s1.Len() > s2.Len() {
		return false
	}
	for k := range s1 {
		if _, ok := s2[k]; !ok {
			return false
		}
	}
	return true
}

// Disjoint returns true iff Set s1 and Set s2 contain no
// equal elements.
func Disjoint(s1, s2 Set) bool {
	for k := range s1 {
		if _, ok := s2[k]; ok {
			return false
		}
	}
	return true
}

// Intersection returns a new Set with every element in both
// Set s1 and Set s2.
func Intersection(s1, s2 Set) Set {
	toReturn := New()
	for k := range s1 {
		if _, ok := s2[k]; ok {
			toReturn.Add(k)
		}
	}
	return toReturn
}

// Union returns a new Set with every element in either s1
// or s2.
func Union(s1, s2 Set) Set {
	toReturn := New()
	for k := range s1 {
		toReturn.Add(k)
	}
	for k := range s2 {
		toReturn.Add(k)
	}
	return toReturn
}

// Difference returns a new Set with all elements of s1 that aren't
// also in s2.
func Difference(s1, s2 Set) Set {
	toReturn := New()
	for k := range s1 {
		if _, ok := s2[k]; !ok {
			toReturn.Add(k)
		}
	}
	return toReturn
}

// SymmetricDifference returns a new Set with elements that
// are in either s1 or s2, but not both
func SymmetricDifference(s1, s2 Set) Set {
	return Union(Difference(s1, s2), Difference(s2, s1))
}
