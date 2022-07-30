package solution

// ----
// stack
type stack []int

func newStack() stack {
	return make([]int, 0, 32)
}

func (s *stack) isEmpty() bool {
	return len((*s)) <= 0
}

func (s *stack) push(node int) {
	(*s) = append((*s), node)
}

// panic on empty stack
func (s *stack) pop() int {
	v := (*s)[len((*s))-1]
	(*s) = (*s)[:len((*s))-1]
	return v
}

// end of stack
// ----
