// Package gigasecond solves Exercism's fouth Go problem.
package gigasecond

import (
	"fmt"
	"math"
	"time"
)

// Constant declaration.
const testVersion = 4

// AddGigasecond adds a gigasecond to a t and returns the result.
func AddGigasecond(t time.Time) time.Time {
	gs := int(math.Pow(10, 9))
	d, err := time.ParseDuration(fmt.Sprintf("%ds", gs))
	if err != nil {
		return time.Now()
	}
	t = t.Add(d)
	return t
}
