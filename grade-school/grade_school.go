// Package school solves an Exercism challenge.
package school

import (
	"sort"
)

// Grade is an integer (grade level) and a slice
// of strings (student names).
type Grade struct {
	level    int
	students []string
}

// School is a slice of grades and a size.
type School []Grade

// New returns a new School.
func New() *School {
	return &School{}
}

// Enrollment returns a slice of grades.
func (s *School) Enrollment() []Grade {
	sort.Sort(s)
	return []Grade(*s)
}

// Add adds a student to the given school.
func (s *School) Add(n string, l int) {
	// Do nothing if l is not a valid grade
	if l < minLevel || l > maxLevel {
		return
	}
	// If we find the grade level, append the student
	// and return
	for i := range *s {
		if (*s)[i].level == l {
			(*s)[i].students = append((*s)[i].students, n)
			sort.Strings((*s)[i].students)
			return
		}
	}
	// We don't have grade level l yet.
	*s = append(*s, Grade{l, []string{n}})
	return
}

// Grade returns the corresponding grade level.
func (s *School) Grade(l int) []string {
	for _, g := range *s {
		if g.level == l {
			return g.students
		}
	}
	return []string{}
}

// Len returns the length of the grades slice.
func (s *School) Len() int {
	return len(*s)
}

// Less returns whether the grade at position
// i should be before the grade at position j.
func (s *School) Less(i, j int) bool {
	if (*s)[i].level < (*s)[j].level {
		return true
	}
	return false
}

// Swap switches the position of grades i and j
// in s.grades.
func (s *School) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}
