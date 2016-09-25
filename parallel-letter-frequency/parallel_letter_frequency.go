// Package letter solves an Exercism challenge.
package letter

// ConcurrentFrequency counts the frequency of various letters
// concurrently.
func ConcurrentFrequency(s []string) FreqMap {
	var results FreqMap
	msg := make(chan FreqMap)

	for i := 0; i < len(s); i++ {
		// i needs to be evaluated before it gets incremented
		// at the top of the loop, so I made this function
		// accept a paramater, and pass i to it (because
		// paramaters to functions are evaluated before
		// the goroutine is spawned). Not doing this caused
		// index out of range errors because there's a race
		// condition on i.
		go func(i int) { msg <- Frequency(s[i]) }(i)
	}

	for i := 0; i < len(s); i++ {
		res := <-msg
		if i == 0 {
			results = res
			continue
		}
		for char, val := range res {
			results[char] += val
		}
	}

	return results
}
