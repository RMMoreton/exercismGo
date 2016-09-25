// Package kindergarten solves an Exercism challenge.
package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

// A Garden is a map from children's names to the plants that those children
// are responsible for.
type Garden map[string][]string

// NewGarden returns a new, filled out garden.
func NewGarden(diagram string, children []string) (*Garden, error) {
	// Normalize the diagram.
	diagramTrimmed := strings.Trim(diagram, "\n")
	if diagramTrimmed == diagram {
		return nil, errors.New("diagram didn't start with a newline")
	}
	rows := strings.Split(diagramTrimmed, "\n")
	if len(rows) != 2 {
		return nil, errors.New("invalid number of rows in diagram")
	}
	if len(rows[0]) != len(rows[1]) {
		return nil, errors.New("diagram row lengths do not match")
	}
	if len(rows[0]) != 2*len(children) {
		return nil, errors.New("diagram rows are not the correct length")
	}
	for i := 0; i < len(rows[0]); i++ {
		c1 := rows[0][i]
		c2 := rows[1][i]
		if c1 != 'G' && c1 != 'C' && c1 != 'R' && c1 != 'V' {
			return nil, errors.New("invalid symbol in diagram")
		}
		if c2 != 'G' && c2 != 'C' && c2 != 'R' && c2 != 'V' {
			return nil, errors.New("invalid symbol in diagram")
		}
	}
	// Make a copy of the input array and sort it.
	childrenCopy := make([]string, len(children))
	copy(childrenCopy, children)
	sort.Strings(childrenCopy)
	// Make sure all children are unique.
	for i := 1; i < len(childrenCopy); i++ {
		if childrenCopy[i-1] == childrenCopy[i] {
			return nil, errors.New("same child name seen more than once")
		}
	}
	// Add each child to the Garden, and their plants.
	g := make(Garden)
	for i, r := range rows[0] {
		childName := childrenCopy[i/2]
		g[childName] = append(g[childName], lookupPlant(r))
	}
	for i, r := range rows[1] {
		childName := childrenCopy[i/2]
		g[childName] = append(g[childName], lookupPlant(r))
	}
	return &g, nil
}

// lookupPlant takes a rune and returns the name of the corresponding plant.
func lookupPlant(r rune) string {
	switch r {
	case 'G':
		return "grass"
	case 'C':
		return "clover"
	case 'R':
		return "radishes"
	case 'V':
		return "violets"
	}
	return "default"
}

// Plants takes a child's name and returns the plants that child is
// responsible for.
func (g *Garden) Plants(child string) ([]string, bool) {
	plants, ok := (*g)[child]
	if !ok {
		return []string{}, false
	}
	return plants, true
}
