// Package meetup solves an Exercism challenge.
package meetup

import (
	"time"
)

// Necessary for testing.
const testVersion = 3

// A Weekschedule is a number that corresponds to "first", "second", etc.
type WeekSchedule int

// These constants define the First, Second, etc. weekschedules.
const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Teenth
	Last
)

// Day returns what day of the month the described day is/was.
func Day(w WeekSchedule, d time.Weekday, m time.Month, y int) int {
	var result int
	var t time.Time
	switch w {
	case First:
		t = time.Date(y, m, 1, 0, 0, 0, 0, time.UTC)
		result = 1
	case Second:
		t = time.Date(y, m, 8, 0, 0, 0, 0, time.UTC)
		result = 8
	case Third:
		t = time.Date(y, m, 15, 0, 0, 0, 0, time.UTC)
		result = 15
	case Fourth:
		t = time.Date(y, m, 22, 0, 0, 0, 0, time.UTC)
		result = 22
	case Teenth:
		t = time.Date(y, m, 13, 0, 0, 0, 0, time.UTC)
		result = 13
	case Last:
		t = time.Date(y, m+1, 1, 0, 0, 0, 0, time.UTC)
		t = t.AddDate(0, 0, -7)
		result = t.Day()
	}
	dayOfWeek := t.Weekday()
	delta := int(d - dayOfWeek)
	if delta < 0 {
		delta += 7
	}
	result += delta
	return result
}
