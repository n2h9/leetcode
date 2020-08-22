package solution

func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	accumulator := make(map[int]map[int]int, n)
	var stack *stackNode

	leftBorderNode := &valueNode{
		iterationsLeft:   k - 1,
		sum:              cardPoints[0],
		leftBorderIndex:  1,
		rightBorderIndex: n - 1,
	}
	rightBorderNode := &valueNode{
		iterationsLeft:   k - 1,
		sum:              cardPoints[1],
		leftBorderIndex:  0,
		rightBorderIndex: n - 2,
	}

	if cardPoints[0] < cardPoints[n-1] {
		stack = stack_push(stack, stackNode_new(leftBorderNode))
		stack = stack_push(stack, stackNode_new(rightBorderNode))
	} else {
		stack = stack_push(stack, stackNode_new(rightBorderNode))
		stack = stack_push(stack, stackNode_new(leftBorderNode))
	}

	max := 0
	for !stack_isEmpty(stack) {
		node := stack_top(stack)
		stack = stack_pop(stack)
		iterationsLeft := node.iterationsLeft
		sum := node.sum
		leftBorderIndex := node.leftBorderIndex
		rightBorderIndex := node.rightBorderIndex

		for iterationsLeft > 0 {
			var moveAssideNode *valueNode
			if cardPoints[leftBorderIndex] < cardPoints[rightBorderIndex] {
				moveAssideNode = &valueNode{
					iterationsLeft:   iterationsLeft - 1,
					sum:              sum + cardPoints[leftBorderIndex],
					leftBorderIndex:  leftBorderIndex + 1,
					rightBorderIndex: rightBorderIndex,
				}
				sum += cardPoints[rightBorderIndex]
				rightBorderIndex--
				iterationsLeft--
			} else {
				moveAssideNode = &valueNode{
					iterationsLeft:   iterationsLeft - 1,
					sum:              sum + cardPoints[rightBorderIndex],
					leftBorderIndex:  leftBorderIndex,
					rightBorderIndex: rightBorderIndex - 1,
				}
				sum += cardPoints[leftBorderIndex]
				leftBorderIndex++
				iterationsLeft--
			}
			stack = stack_push(stack, stackNode_new(moveAssideNode))
		}

		// if v, ok := accumulator[node.leftBorderIndex]; ok {
		// 	if v2, ok := v[node.rightBorderIndex]; ok {
		// 		// if v2 > max {
		// 		// 	max = v2
		// 		// }
		// 		continue
		// 	}
		// } else {
		// 	accumulator[node.leftBorderIndex] = make(map[int]int, 1000)
		// }
	}

	return max
}

type valueNode struct {
	iterationsLeft   int
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
