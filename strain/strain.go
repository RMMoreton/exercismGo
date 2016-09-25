// Package strain solves an Exercism challenge.
package strain

// Ints is just a slice of ints.
type Ints []int

// Lists is a slice of slices of ints.
type Lists [][]int

// Strings is a slice of strings.
type Strings []string

// Keep returns a new Ints containing only those elements
// v of i such that f(v) == true.
func (i Ints) Keep(f func(int) bool) Ints {
	var toReturn Ints
	for _, v := range i {
		if f(v) {
			toReturn = append(toReturn, v)
		}
	}
	return toReturn
}

// Discard returns a new Ints containing only those elements
// v of i such that f(v) == false.
func (i Ints) Discard(f func(int) bool) Ints {
	var toReturn Ints
	for _, v := range i {
		if !f(v) {
			toReturn = append(toReturn, v)
		}
	}
	return toReturn
}

// Keep returns a new Lists containing only those elements
// v of l such that f(v) == true.
func (l Lists) Keep(f func([]int) bool) Lists {
	var toReturn Lists = make([][]int, 0, len(l))
	for _, v := range l {
		if f(v) {
			toReturn = append(toReturn, v)
		}
	}
	return toReturn
}

// Keep returns a new Strings containing only those elements
// v of s such that f(v) == true.
func (s Strings) Keep(f func(string) bool) Strings {
	var toReturn Strings = make([]string, 0, len(s))
	for _, v := range s {
		if f(v) {
			toReturn = append(toReturn, v)
		}
	}
	return toReturn
}
