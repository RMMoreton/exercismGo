// Package leap is exercism's second Go challenge.
package leap

// testVersion should match the targetTestVersion in the test file.
const testVersion = 2

// IsLeapYear passes a year and returns whether that year is a leap year.
func IsLeapYear(y int) bool {
	if y%400 == 0 {
		return true
	}
	if y%100 == 0 {
		return false
	}
	if y%4 == 0 {
		return true
	}
	return false
}
