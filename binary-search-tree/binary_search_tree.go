// Package binarysearchtree solves an Exercism challenge.
package binarysearchtree

import ()

// A SearchTreeData is a binary search tree.
type SearchTreeData struct {
	data  int
	left  *SearchTreeData
	right *SearchTreeData
}

// Bst returns a new SearchTreeData with a single node, with that node's
// data set to x.
func Bst(x int) *SearchTreeData {
	return &SearchTreeData{data: x}
}

// Insert inserts data value x into the binary search tree.
func (t *SearchTreeData) Insert(x int) {
	for {
		if x <= t.data {
			if t.left != nil {
				t = t.left
				continue
			}
			t.left = Bst(x)
			return
		} else {
			if t.right != nil {
				t = t.right
				continue
			}
			t.right = Bst(x)
			return
		}
	}
}

// MapString maps a function onto the elements of t, in order, and
// returns the slice of the results.
func (t *SearchTreeData) MapString(f func(int) string) []string {
	res := make([]string, 0)
	if t.left != nil {
		res = append(res, t.left.MapString(f)...)
	}
	res = append(res, f(t.data))
	if t.right != nil {
		res = append(res, t.right.MapString(f)...)
	}
	return res
}

// MapInt maps a function onto the elements of t, in order, and returns
// the slice of the results.
func (t *SearchTreeData) MapInt(f func(int) int) []int {
	res := make([]int, 0)
	if t.left != nil {
		res = append(res, t.left.MapInt(f)...)
	}
	res = append(res, f(t.data))
	if t.right != nil {
		res = append(res, t.right.MapInt(f)...)
	}
	return res
}
