// Package react solves an Exercism challenge.
package react

// testVersion is required for testing.
const testVersion = 4

// A reactor implements the Reactor interface.
type reactor struct {
	dependencies    map[Cell][]Cell      // the map of cell dependencies in this reactor
	neededCallbacks map[Cell][]func(int) // a map of cells to callback functions
}

// An inCell is an input cell.
type inCell struct {
	val int      // value of the cell
	re  *reactor // the reactor to which this cell belongs
}

// A comCell is a compute cell.
type comCell struct {
	f          func(int, int) int           // function to compute this cell's value
	val        int                          // this cell's value
	callbacks  map[CallbackHandle]func(int) // all callbacks on this cell
	in1        Cell                         // one cell from which this cell derives it's value
	in2        Cell                         // another cell form which this cell derives it's value
	callHandle int                          // used to give out unique CallbackHandles
	re         *reactor                     // the reactor to which this cell is attached
}

// New returns a new reator.
func New() *reactor {
	return &reactor{
		dependencies:    make(map[Cell][]Cell),
		neededCallbacks: make(map[Cell][]func(int)),
	}
}

// CreateInput returns a new input cell with value n.
func (r *reactor) CreateInput(n int) InputCell {
	return &inCell{
		val: n,
		re:  r,
	}
}

// CreateCompute1 returns a new compute-1 cell.
func (r *reactor) CreateCompute1(c Cell, f func(int) int) ComputeCell {
	return r.CreateCompute2(c, new(inCell), func(x, y int) int { return f(x) })
}

// CreateCompute2 returns a new compute-2 cell.
func (r *reactor) CreateCompute2(c1, c2 Cell, f func(int, int) int) ComputeCell {
	toReturn := &comCell{
		f:          f,
		val:        f(c1.Value(), c2.Value()),
		callbacks:  make(map[CallbackHandle]func(int)),
		in1:        c1,
		in2:        c2,
		callHandle: 0,
		re:         r,
	}
	// Add c1 to the dependency map, and add toReturn to c1's dependents.
	if _, ok := r.dependencies[c1]; !ok {
		r.dependencies[c1] = []Cell{toReturn}
	} else {
		r.dependencies[c1] = append(r.dependencies[c1], toReturn)
	}
	// Add c2 to the dependency map, and add toReturn to c2's dependents.
	if _, ok := r.dependencies[c2]; !ok {
		r.dependencies[c2] = []Cell{toReturn}
	} else {
		r.dependencies[c2] = append(r.dependencies[c2], toReturn)
	}

	return toReturn
}

// SetValue allows us to set the value of an input cell.
func (c *inCell) SetValue(n int) {
	if c.val != n {
		c.val = n
		if propList, ok := c.re.dependencies[c]; ok {
			for _, subCell := range propList {
				subCell.Value() // propogate the change
			}
		}
		// Call all needed callbacks
		for c, callbacks := range c.re.neededCallbacks {
			for _, cb := range callbacks {
				cb(c.Value())
			}
		}
		// Clear the needed callbacks
		c.re.neededCallbacks = make(map[Cell][]func(int))
	}
}

// Value returns the value of the input cell.
func (c *inCell) Value() int {
	return c.val
}

// Value recomputes c's value, and propogates any changes.
// Cell c is then added to a set of cells whose callbacks will
// be called by the cell which initiated this propogation
// of changes.
func (c *comCell) Value() int {
	oldVal := c.val
	newVal := c.f(c.in1.Value(), c.in2.Value())
	if newVal != oldVal {
		c.val = newVal
		if propList, ok := c.re.dependencies[c]; ok {
			for _, subCell := range propList {
				subCell.Value() // propogate this change
			}
		}
		// Add c's callbacks to the neededCallback map (only if they're not already there)
		if _, ok := c.re.neededCallbacks[c]; !ok {
			c.re.neededCallbacks[c] = make([]func(int), 0)
			for _, cb := range c.callbacks {
				c.re.neededCallbacks[c] = append(c.re.neededCallbacks[c], cb)
			}
		}
	}
	return c.val
}

// AddCallback adds a callback function to the compute-2 cell,
// and returns a handle by which that function may be indentified.
func (c *comCell) AddCallback(f func(int)) CallbackHandle {
	c.callbacks[c.callHandle] = f
	c.callHandle++
	return c.callHandle - 1
}

// RemoveCallback removes a callback function from a compute-2 cell.
func (c *comCell) RemoveCallback(h CallbackHandle) {
	delete(c.callbacks, h)
}
