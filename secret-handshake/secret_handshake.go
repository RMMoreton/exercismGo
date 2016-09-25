// Package secret solves an Exercism challenge.
package secret

// In my first iteration, I was declaring constants and using those
// constants to index into a slice of strings. Skipping the constants
// and just making the string slice is way more readable, though. Thanks
// to @LukasMa for that.
var code = []string{
	"wink",
	"double blink",
	"close your eyes",
	"jump",
}

// Handshake takes an integer and returns the corresponding secret handshake.
func Handshake(x int) []string {
	if x < 0 { // x must be non-negative.
		return nil
	}
	shake := make([]string, 0)
	for i, s := range code {
		if x&(1<<uint(i)) != 0 {
			shake = append(shake, s)
		}
	}
	if x&(1<<len(code)) != 0 { // Reverse if need be.
		for i, j := 0, len(shake)-1; i < j; i, j = i+1, j-1 {
			shake[i], shake[j] = shake[j], shake[i]
		}
	}
	return shake
}
