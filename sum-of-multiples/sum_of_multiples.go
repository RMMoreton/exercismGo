// Package summultiples solves an Exercism challenge.
package summultiples

// MultipleSummer returns a function that takes an n
// and sums all multiples of mults up to n.
func MultipleSummer(mults ...int) func(int) int {
	return func(n int) int {
		var res int
		for i := 0; i < len(mults); i++ {
			for j := mults[i]; j < n; j += mults[i] {
				present := false
				// Check if I've added j before
				for k := 0; k < i; k++ {
					if j%mults[k] == 0 {
						present = true
					}
				}
				if !present {
					res += j
				}
			}
		}
		return res
	}
}
