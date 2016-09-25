// Package matrix solves an Exercism challenge.
package matrix

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// A Matrix is exactly what it sounds like.
type Matrix [][]int

// Rows returns the rows of a matrix by copying the rows of matrix
// m into a new slice, and returning that slice.
func (m *Matrix) Rows() [][]int {
	numRows := len(*m)
	numCols := len((*m)[0])
	r := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		r[i] = make([]int, numCols)
		for j := 0; j < numCols; j++ {
			r[i][j] = (*m)[i][j]
		}
	}
	return r
}

// Columns returns the columns of a matrix by copying the columns of matrix
// m into the rows of a new slice, and returning that slice.
func (m *Matrix) Cols() [][]int {
	numRows := len(*m)
	numCols := len((*m)[0])
	c := make([][]int, numCols)
	for i := 0; i < numCols; i++ {
		c[i] = make([]int, numRows)
		for j := 0; j < numRows; j++ {
			c[i][j] = (*m)[j][i]
		}
	}
	return c
}

// Set sets the value at position (r, c) in matrix m to value val.
func (mA *Matrix) Set(r, c, val int) bool {
	m := *mA
	if r < 0 || r >= len(m) {
		return false
	}
	if c < 0 || c >= len(m[0]) {
		return false
	}
	m[r][c] = val
	return true
}

// New creates a new matrix from a string.
func New(s string) (*Matrix, error) {
	m := *new(Matrix)
	rows := strings.Split(s, "\n")
	if nil == rows {
		return nil, errors.New("No rows found")
	}
	var numCols int
	for i := 0; i < len(rows); i++ {
		cols := strings.Split(strings.Trim(rows[i], " "), " ")
		// Do some checking to make sure all columns have the same length.
		if i == 0 {
			numCols = len(cols)
		} else {
			if len(cols) != numCols {
				return nil, errors.New("Mismatched number of columns")
			}
		}
		// Put everything into m
		m = append(m, make([]int, 0))
		for j := 0; j < len(cols); j++ {
			num, err := strconv.Atoi(cols[j])
			if nil != err {
				return nil, errors.New(fmt.Sprintf("Non-number found in input: %s", cols[j]))
			}
			m[i] = append(m[i], num)
		}
	}
	return &m, nil
}
