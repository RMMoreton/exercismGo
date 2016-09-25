// Package slice solves an Exercism challenge.
package slice

// All returns all substrings of length n in s.
func All(n int, s string) []string {
	if n < 1 || n > len(s) {
		return []string{}
	}
	var result = make([]string, 0, len(s)-n+1)
	for start, end := 0, n; end <= len(s); start, end = start+1, end+1 {
		result = append(result, s[start:end])
	}
	return result
}

// UnsafeFirst returns the first substring of length n in s.
func UnsafeFirst(n int, s string) string {
	if n < 1 || n > len(s) {
		return ""
	}
	return s[0:n]
}

// First returns the first substring of length n in s and true if
// such a substring exists, or the empty string and false if not.
func First(n int, s string) (string, bool) {
	if n < 1 || n > len(s) {
		return "", false
	}
	return s[0:n], true
}
