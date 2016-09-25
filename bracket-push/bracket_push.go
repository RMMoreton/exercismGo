// Package brackets solves an Exercism challenge.
package brackets

// testVersion is required for the testing suite
const testVersion = 3

// Stack has the three basic stack methods attached to it.
type Stack []rune

// Bracket returns true iff the string is a correctly
// bracketed expression.
func Bracket(s string) (bool, error) {
	var matches = map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
	}
	var stack Stack
	for _, r := range s {
		if r == '{' || r == '(' || r == '[' { // r is an opening bracket
			stack.Push(r)
			continue
		}
		top, ok := stack.Peek() // r is a closing bracket
		switch {
		case !ok: // stack's empty, so r has nothing to close
			return false, nil
		case r == matches[top]: // r closes the top element of the stack
			stack.Pop()
		default: // stack's not empty, but r doesn't close the top element
			return false, nil
		}
	}
	return len(stack) == 0, nil // every bracket needs to have been matched
}

// Push appends the passed rune onto the end of the stack.
func (s *Stack) Push(r rune) {
	*s = append(*s, r)
}

// Peek returns the top rune on the stack, or a false boolean if
// the stack is empty.
func (s *Stack) Peek() (rune, bool) {
	var r rune
	if len(*s) == 0 {
		return r, false
	}
	r = (*s)[len(*s)-1]
	return r, true
}

// Pop removes and returns the top element of the stack if possible,
// or returns a false boolean if the stack is empty.
func (s *Stack) Pop() (rune, bool) {
	r, ok := s.Peek()
	if !ok {
		return r, false
	}
	*s = (*s)[:len(*s)-1]
	return r, true
}
