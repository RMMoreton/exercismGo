// Package hello implements the solution to Exercism's 'hello world' test.
package hello

import (
	"fmt"
)

const testVersion = 2

// HelloWorld passes a name, which may be empty, and returns a simple greeting.
func HelloWorld(name string) string {
	if name == "" {
		name = "World"
	}
	r := fmt.Sprintf("Hello, %s!", name)
	return r
}