// Package connect solves an Exercism challenge.
package connect

// testVersion is required for testing.
const testVersion = 2

// A pos is an y,x coordinate on the board.
type pos [2]int

// ResultOf finds the winner of a game of connect.
func ResultOf(board []string) (string, error) {
	// oWinCond reports whether an O at position p gives 'O' the win.
	oWinCond := func(p pos) bool {
		if p[0] == len(board)-1 {
			return true
		}
		return false
	}

	// xWinCond reports whether an X at position p gives 'X' the win.
	xWinCond := func(p pos) bool {
		if p[1] == len(board[0])-1 {
			return true
		}
		return false
	}

	// search takes an initial queue and does a BFS, checking winCond for
	// every processed node.
	search := func(q []pos, winCond func(pos) bool, c byte) bool {
		if len(q) == 0 {
			return false
		}
		used := make([][]byte, len(board))
		for i := 0; i < len(used); i++ {
			used[i] = make([]byte, len(board[0]))
		}

		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			y, x := cur[0], cur[1]
			if y < 0 || y >= len(board) || x < 0 || x >= len(board[0]) {
				continue
			}
			if used[y][x] == 1 {
				continue
			}
			if board[y][x] != c {
				continue
			}
			if winCond(cur) {
				return true
			}
			used[y][x] = 1
			neighbors := []pos{{y - 1, x}, {y - 1, x + 1}, {y, x - 1}, {y, x + 1}, {y + 1, x - 1}, {y + 1, x}}
			q = append(q, neighbors...)
		}
		return false
	}

	// See if 'o' wins.
	oStartQ := make([]pos, len(board[0]))
	for i := 0; i < len(oStartQ); i++ {
		oStartQ[i] = pos{0, i}
	}
	oWin := search(oStartQ, oWinCond, 'O')

	// See if 'x' wins.
	xStartQ := make([]pos, len(board))
	for i := 0; i < len(board); i++ {
		xStartQ[i] = pos{i, 0}
	}
	xWin := search(xStartQ, xWinCond, 'X')

	// Return the correct result.
	switch {
	case oWin && xWin:
		return "", nil
	case oWin:
		return "O", nil
	case xWin:
		return "X", nil
	default:
		return "", nil
	}
}
