// Package robot solves an Exercism challenge.
package robot

import (
	"fmt"
)

/*
Step 1
*/

// Directions.
const (
	N = iota
	E
	S
	W
)

// Right turns the robot right.
func Right() {
	Facing += 1
	if Facing > 3 {
		Facing = 0
	}
}

// Left turns the robot left.
func Left() {
	Facing -= 1
	if Facing < 0 {
		Facing = 3
	}
}

// Advance advances the robot one unit in the direction it's facing.
func Advance() {
	switch Facing {
	case N:
		Y += 1
	case E:
		X += 1
	case S:
		Y -= 1
	case W:
		X -= 1
	}
}

// String returns the string representation of a direction.
func (d Dir) String() string {
	switch d {
	case N:
		return "North"
	case E:
		return "East"
	case S:
		return "South"
	case W:
		return "West"
	}
	return ""
}

/*
Step 2
*/

// An Action is one of 'L', 'R', or 'A'.
type Action byte

// Robot simulates a robot in a room.
func Robot(cmd chan Command, act chan Action) {
	for c := range cmd {
		act <- Action(c)
	}
	close(act)
}

// Room simulates the room my robot is in.
func Room(extent Rect, cur DirAt, act chan Action, rep chan DirAt) {
	// Execute the actions sent by the robot.
	for a := range act {
		switch a {
		case 'A':
			cur = advance(cur, extent)
		default:
			cur.Dir = turn(cur.Dir, a)
		}
	}
	rep <- cur
	close(rep)
}

// turn takes a facing direction and a turn direction, and returns a new
// direction.
func turn(facing Dir, turn Action) Dir {
	switch turn {
	case 'R':
		facing += 1
		if facing > 3 {
			facing = 0
		}
	case 'L':
		facing -= 1
		if facing < 0 {
			facing = 3
		}
	}
	return facing
}

// advance takes a DirAt and an extent, and attempts to advance the DirAt.
// A new DirAt is returned.
func advance(cur DirAt, extent Rect) DirAt {
	switch cur.Dir {
	case N:
		if cur.Northing+1 <= extent.Max.Northing {
			cur.Northing += 1
		}
	case E:
		if cur.Easting+1 <= extent.Max.Easting {
			cur.Easting += 1
		}
	case S:
		if cur.Northing-1 >= extent.Min.Northing {
			cur.Northing -= 1
		}
	case W:
		if cur.Easting-1 >= extent.Min.Easting {
			cur.Easting -= 1
		}
	}
	return cur
}

/*
Step 3
*/

// An Action3 is an action in step 3.
type Action3 struct {
	name string
	a    Action
}

// Robot3 simulates a single robot in step 3.
func Robot3(name, script string, action chan Action3, log chan string) {
	if name == "" {
		log <- "unnamed robot"
		return
	}
	for _, c := range []byte(script) {
		if c != 'R' && c != 'L' && c != 'A' {
			log <- fmt.Sprintf("%s had bad command %q", name, c)
			action <- Action3{name, 'S'}
			return
		}
		action <- Action3{name, Action(c)}
	}
	action <- Action3{name, 'S'}
	return
}

// Room3 simulates a room in step 3.
func Room3(extent Rect, robots []Place, action chan Action3, report chan []Place, log chan string) {
	// advance3 checks what will happen if the passed robot advances,
	// and logs any problems. If there are none, it advances the robot.
	advance3 := func(r Place) Place {
		temp := advance(r.DirAt, extent)
		if temp == r.DirAt {
			log <- fmt.Sprintf("%s tried to move into a wall", r.Name)
			return r
		}
		for _, r2 := range robots {
			if temp.Northing == r2.Northing && temp.Easting == r2.Easting {
				log <- fmt.Sprintf("%s tried to move into %s", r.Name, r2.Name)
				return r
			}
		}
		return Place{r.Name, temp}
	}

	// Make sure all robots start in good starting positions.
	for i := 0; i < len(robots); i++ {
		r1 := robots[i]
		// Make sure r1 is in the room.
		switch {
		case r1.Northing < extent.Min.Northing || r1.Easting < extent.Min.Easting:
			log <- fmt.Sprintf("%s outside room", r1.Name)
			report <- robots
			return
		case r1.Northing > extent.Max.Northing || r1.Easting > extent.Max.Easting:
			log <- fmt.Sprintf("%s outside room", r1.Name)
			report <- robots
			return
		}
		for j := i + 1; j < len(robots); j++ {
			r2 := robots[j]
			// Make sure r1 and r2 are at different positions.
			if r1.Northing == r2.Northing && r1.Easting == r2.Easting {
				log <- fmt.Sprintf("%s and %s placed at same position", r1.Name, r2.Name)
				report <- robots
				return
			}
		}
	}

	// Make a map of robot string names to their indexes in robots.
	numToTerminate := len(robots)
	rMap := make(map[string]int)
	for i, r := range robots {
		if r.Name == "" {
			numToTerminate--
			continue
		}
		if _, ok := rMap[r.Name]; ok {
			log <- fmt.Sprintf("name %s used twice", r.Name)
			report <- robots
			return
		}
		rMap[r.Name] = i
	}

	// Get actions from robots until they've all reported 'S'.
	for numToTerminate > 0 {
		a := <-action
		// Log unknown names and terminate.
		if _, ok := rMap[a.name]; !ok && a.name != "" {
			log <- fmt.Sprintf("unknown name %s", a.name)
			report <- robots
			return
		}
		i := rMap[a.name]
		p := robots[i]
		switch a.a {
		case 'A':
			p = advance3(p)
		case 'S':
			numToTerminate--
		default:
			p.Dir = turn(p.Dir, a.a)
		}
		robots[i] = p
	}

	report <- robots
	return
}
