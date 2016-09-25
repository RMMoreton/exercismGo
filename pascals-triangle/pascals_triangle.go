// Package pascal solves an Exercism challenge.
package pascal

// Triangle returns an 2D slice of ints holding Pascals Triangle.
func Triangle(n int) [][]int {
	if n <= 0 { // n must be at least 1.
		return nil
	}
	var result = make([][]int, n)
	result[0] = make([]int, 1)
	result[0][0] = 1
	for i := 1; i < n; i++ {
		result[i] = make([]int, i+1)
		result[i][0] = 1
		// left and right index into the upper level of the triangle; loop
		// until right isn't a valid index.
		for left, right := 0, 1; right < i; left, right = left+1, right+1 {
			result[i][right] = result[i-1][left] + result[i-1][right]
		}
		result[i][i] = 1
	}
	return result
}
