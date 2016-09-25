// Package binarysearch solves an Exercism challenge.
package binarysearch

import (
	"fmt"
)

// SearchInts searches for an element in a slice of ints.
func SearchInts(a []int, x int) int {
	high := len(a) - 1
	low := 0
	for high >= low {
		mid := (high + low) / 2
		if x == a[mid] {
			for mid-1 >= 0 && a[mid-1] == x {
				mid = mid - 1
			}
			return mid
		} else if x < a[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return high + 1
}

// Message returns a message based on where x is in a.
func Message(a []int, x int) string {
	i := SearchInts(a, x)
	switch {
	case len(a) == 0:
		return "slice has no values"
	case i == len(a):
		return fmt.Sprintf("%d > all %d values", x, len(a))
	case a[i] == x:
		switch {
		case i == 0:
			return fmt.Sprintf("%d found at beginning of slice", x)
		case i == len(a)-1:
			return fmt.Sprintf("%d found at end of slice", x)
		default:
			return fmt.Sprintf("%d found at index %d", x, i)
		}
	case i == 0:
		return fmt.Sprintf("%d < all values", x)
	default:
		return fmt.Sprintf("%d > %d at index %d, < %d at index %d", x, a[i-1], i-1, a[i], i)
	}
}
