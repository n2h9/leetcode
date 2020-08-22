package solution

func validateStackSequences(pushed []int, popped []int) bool {
	n := len(pushed)
	if n <= 0 {
		return true
	}
	stack := newStack(n)
	stack.push(pushed[0])
	pui, poi := 1, 0

	for {
		if !stack.isEmpty() && poi < n && stack.top() == popped[poi] {
			stack.pop()
			poi++
			continue
		}

		if pui < n {
			stack.push(pushed[pui])
			pui++
			continue
		}

		break
	}

	return stack.isEmpty()
}

type stack struct {
	topIndex int
	storage  []int
}

func newStack(len int) *stack {
	return &stack{
		topIndex: -1,
		storage:  make([]int, len),
	}
}

func (s *stack) isEmpty() bool {
	return s.topIndex < 0
}

func (s *stack) push(value int) {
	s.topIndex++
	s.storage[s.topIndex] = value
}

func (s *stack) pop() {
	s.topIndex--
}

func (s *stack) top() int {
	return s.storage[s.topIndex]
}
