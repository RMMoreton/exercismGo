// Package wordsearch solves an Exercism challenge.
package wordsearch

// Required for testing.
const testVersion = 2

// Solve takes a slice of words and a puzzle and finds the words in the
// puzzle.
func Solve(words, puzzle []string) (map[string][2][2]int, error) {
	// search looks for word w starting at position (i, j).
	search := func(w string, i, j int) (bool, int, int) {
		for dI := -1; dI < 2; dI++ {
			for dJ := -1; dJ < 2; dJ++ {
				// If they're both zero, we're not searching for anything.
				if dI == 0 && dJ == 0 {
					continue
				}
				// Copy i and j so I can mess with them, and make a word index.
				tI, tJ := i, j
				wI := 0
				for ; wI < len(w); wI, tI, tJ = wI+1, tI+dI, tJ+dJ {
					// Make sure tI and tJ are valid puzzle positions.
					if tI < 0 || tI >= len(puzzle) || tJ < 0 || tJ >= len(puzzle[0]) {
						break
					}
					if w[wI] != puzzle[tI][tJ] {
						break
					}
				}
				if wI == len(w) {
					return true, tI - dI, tJ - dJ
				}
			}
		}
		return false, 0, 0
	}

	// Make the map I'll return, and loop through the words searching for them.
	ret := make(map[string][2][2]int)
	for _, w := range words {
		var sI, sJ, eI, eJ int
		found := false
		for i := 0; i < len(puzzle) && !found; i++ {
			for j := 0; j < len(puzzle[0]) && !found; j++ {
				sI, sJ = i, j
				found, eI, eJ = search(w, i, j)
			}
		}
		if found {
			ret[w] = [2][2]int{[2]int{sJ, sI}, [2]int{eJ, eI}}
		} else {
			ret[w] = [2][2]int{[2]int{-1, -1}, [2]int{-1, -1}}
		}
	}
	return ret, nil
}
