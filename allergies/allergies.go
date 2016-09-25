// Package allergies solves an Exercism challenge.
package allergies

// Possible is an array containing all possible allergens,
// ordered by their alergy score.
var Possible = [8]string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

// Allergies returns a slice of strings of everything
// a person with allergy score n is alergic to.
func Allergies(n int) []string {
	var flag int = 1
	res := make([]string, 0, 8)
	for _, s := range Possible {
		if n&flag > 0 {
			res = append(res, s)
		}
		flag = flag << 1
	}
	return res
}

// AllergicTo returns true if someone with allergy score
// n is allergic to substance s, false otherwise.
func AllergicTo(n int, a string) bool {
	var flag int = 1
	for _, s := range Possible {
		if s == a && n&flag > 0 {
			return true
		}
		flag = flag << 1
	}
	return false
}
