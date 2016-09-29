// Package tournament solves an Exercism challenge.
package tournament

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

// Required for testing.
const testVersion = 3

// LL is the default line length.
const LL = 128

// A record keeps track of the wins, draws, and losses of a single team.
type record struct {
	n          string
	w, d, l, p int
}

// A recordSlice is a slice of records.
type recordSlice []*record

// Tally reads results of football games from r and calculates scores for each
// team, then writes the results to w. If there's bad input, Tally returns a
// non-nil error.
func Tally(r io.Reader, w io.Writer) error {
	// Read all the input.
	teams := make(map[string]*record)
	for {
		l, err := readLine(r)
		if err != nil {
			break
		}
		fields := bytes.Split(l, []byte{';'})
		if len(fields) != 3 {
			continue
		}
		res := string(fields[2])
		if res != "win" && res != "draw" && res != "loss" {
			continue
		}
		t1 := string(fields[0])
		t2 := string(fields[1])
		if _, ok := teams[t1]; !ok {
			teams[t1] = &record{t1, 0, 0, 0, 0}
		}
		if _, ok := teams[t2]; !ok {
			teams[t2] = &record{t2, 0, 0, 0, 0}
		}
		teams[t1].p++
		teams[t2].p++
		switch string(fields[2]) {
		case "win":
			teams[t1].w++
			teams[t2].l++
		case "loss":
			teams[t1].l++
			teams[t2].w++
		case "draw":
			teams[t1].d++
			teams[t2].d++
		}
	}

	// If no valid teams were found, return an error.
	if len(teams) == 0 {
		return errors.New("no valid lines found")
	}

	// Make a slice of the teams and sort it.
	teamSlice := make([]*record, len(teams))
	i := 0
	for _, t := range teams {
		teamSlice[i] = t
		i++
	}
	sort.Sort(recordSlice(teamSlice))

	// Write the results to the passed writer.
	w.Write([]byte(fmt.Sprintf("%-30s | MP |  W |  D |  L |  P\n", "Team")))
	for _, r := range teamSlice {
		w.Write([]byte(fmt.Sprintf("%-30s | %2d | %2d | %2d | %2d | %2d\n", r.n, r.p, r.w, r.d, r.l, 3*r.w+r.d)))
	}

	return nil
}

// readLine reads a line of input from r and returns it.
func readLine(r io.Reader) ([]byte, error) {
	ret := make([]byte, LL)
	single := make([]byte, 1)
	var n, i int
	var err error
	for {
		n, err = r.Read(single)
		if n == 0 {
			break
		}
		if single[0] == '\n' {
			break
		}
		if i == len(ret) {
			tmp := make([]byte, 2*len(ret))
			copy(tmp, ret)
			ret = tmp
		}
		ret[i] = single[0]
		i++
	}
	return ret[:i], err
}

// Len returns the length of a []*record.
func (slice recordSlice) Len() int {
	return len(slice)
}

// Less reports whether the element at i should come before the element at j.
func (slice recordSlice) Less(i, j int) bool {
	iP := 3*slice[i].w + slice[i].d
	jP := 3*slice[j].w + slice[j].d
	switch {
	case iP < jP:
		return false
	case iP == jP:
		if strings.Compare(slice[i].n, slice[j].n) == 1 {
			return false
		}
		return true
	case iP > jP:
		return true
	}
	return false
}

// Swap swaps two records in a slice.
func (slice recordSlice) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
