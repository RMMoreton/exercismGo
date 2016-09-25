// Package matrix solves an Exercism challenge.
package matrix

// A Pair is an ordered set of two integers.
type Pair [2]int

// Saddle searches for saddle points in a given matrix.
func (mA *Matrix) Saddle() []Pair {
	rows := mA.Rows()
	cols := mA.Cols()
	rowMax := make([]int, len(rows))
	colMin := make([]int, len(cols))
	result := make([]Pair, 0)
	for i := 0; i < len(rows); i++ {
		rowMax[i] = findMax(rows[i])
	}
	for i := 0; i < len(cols); i++ {
		colMin[i] = findMin(cols[i])
	}
	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(cols); j++ {
			if rows[i][j] == rowMax[i] && rows[i][j] == colMin[j] {
				result = append(result, Pair{i, j})
			}
		}
	}
	return result
}

// findMax returns the maximum value from a slice of ints.
func findMax(a []int) int {
	max := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	return max
}

// findMin returns the minimum value from a slice of ints.
func findMin(a []int) int {
	min := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] < min {
			min = a[i]
		}
	}
	return min
}
