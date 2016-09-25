// Package clock solves Exercism's third Go exercise.
package clock

import (
	"fmt"
)

// The value of testVersion here must match `targetTestVersion` in the file
// clock_test.go.
const testVersion = 4

// Clock stores hours and minutes, but not days, months, etc.
type Clock struct {
	h, m int
}

// New returns a new clock, after standardizing the hours and minutes.
func New(hour, minute int) Clock {
	var rollMin int = minute / 60
	if minute = minute % 60; minute < 0 {
		minute += 60
		hour -= 1
	}
	if hour = (hour + rollMin) % 24; hour < 0 {
		hour += 24
	}
	return Clock{hour, minute}
}

// String returns a string representation of a clock.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

// Add returns a new clock with 'minutes' added.
func (c Clock) Add(minutes int) Clock {
	return New(c.h, c.m+minutes)
}
