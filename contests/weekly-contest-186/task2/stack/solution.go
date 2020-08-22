package solution

func maxScore(cardPoints []int, k int) int {
	max := 0

	n := len(cardPoints)

	stack := stackNode_new(
		&valueNode{
			iterationLeft:    k - 1,
			sum:              cardPoints[0],
			leftBorderIndex:  1,
			rightBorderIndex: n - 1,
		},
	)
	stack = stack_push(
		stack,
		stackNode_new(
			&valueNode{
				iterationLeft:    k - 1,
				sum:              cardPoints[n-1],
				leftBorderIndex:  0,
				rightBorderIndex: n - 2,
			},
		),
	)

	for !stack_isEmpty(stack) {
		node := stack_top(stack)
		stack = stack_pop(stack)

		if node.sum > max {
			max = node.sum
		}
		if node.iterationLeft > 0 {
			stack = stack_push(
				stack,
				stackNode_new(
					&valueNode{
						iterationLeft:    node.iterationLeft - 1,
						sum:              node.sum + cardPoints[node.leftBorderIndex],
						leftBorderIndex:  node.leftBorderIndex + 1,
						rightBorderIndex: node.rightBorderIndex,
					},
				),
			)
			stack = stack_push(
				stack,
				stackNode_new(
					&valueNode{
						iterationLeft:    node.iterationLeft - 1,
						sum:              node.sum + cardPoints[node.rightBorderIndex],
						leftBorderIndex:  node.leftBorderIndex,
						rightBorderIndex: node.rightBorderIndex - 1,
					},
				),
			)
		}
	}

	return max
}

type valueNode struct {
	iterationLeft    int
	sum              int
	leftBorderIndex  int
	rightBorderIndex int
}

type stackNode struct {
	value *valueNode
	next  *stackNode
}

func stackNode_new(val *valueNode) *stackNode {
	return &stackNode{
		value: val,
	}
}

func stack_push(stack *stackNode, node *stackNode) *stackNode {
	node.next = stack
	return node
}

func stack_pop(stack *stackNode) *stackNode {
	next := stack.next
	stack.next = nil
	return next
}

func stack_top(stack *stackNode) *valueNode {
	return stack.value
}

func stack_isEmpty(stack *stackNode) bool {
	return stack == nil
}
