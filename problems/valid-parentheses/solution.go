package solution

func isValid(s string) bool {
	var stack *node
	for i := range s {
		if isOpen(s[i]) {
			stack = push(stack, &node{Val: s[i]})
			continue
		}
		if stack == nil {
			// close item on empty stack
			return false
		}

		if stack.Val != getAppropriateOpenValue(s[i]) {
			// parentheses of different type
			return false
		}

		// all ok, release top stack
		stack = pop(stack)
	}

	// true if and only if stack is empty at the end
	return stack == nil
}

func isOpen(b byte) bool {
	return b == '(' || b == '{' || b == '['
}

func getAppropriateOpenValue(b byte) byte {
	if b == ')' {
		return '('
	}
	if b == '}' {
		return '{'
	}
	return '['
}

type node struct {
	Val  byte
	Next *node
}

func push(stack *node, n *node) *node {
	n.Next = stack
	return n
}

func pop(stack *node) *node {
	top := stack
	stack = stack.Next
	top.Next = nil

	return stack
}
