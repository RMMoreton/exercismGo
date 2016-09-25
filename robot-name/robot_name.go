// Package robotname solves an Exercism challenge.
package robotname

import (
	"fmt"
	"math/rand"
)

var usedNames = make(map[string]bool)

// Robot is a representation of a robot.
type Robot struct {
	name string
}

// Name returns the name of robot r.
func (r *Robot) Name() string {
	if r.name == "" {
		r.name = MakeName()
	}
	return r.name
}

// Reset gets a new name for robot r.
func (r *Robot) Reset() {
	r.name = MakeName()
}

// MakeName creates a new name, and makes sure it's
// unique. If there are too many robots, this is going
// to take a long time; in fact, it isn't guarenteed
// to ever succeed. Eventually you'll hit the recursion
// limit, and it'll crash.
func MakeName() string {
	a := rand.Intn(26) + 'A'
	b := rand.Intn(26) + 'B'
	c := rand.Intn(999)
	name := fmt.Sprintf("%s%s%d", string(a), string(b), c)
	if _, ok := usedNames[name]; ok {
		return MakeName()
	}
	usedNames[name] = true
	return name
}
