package solution

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	isProcessed := make(map[*TreeNode]bool, 10000)
	stack := push(
		nil,
		&StackNode{
			Val: root,
		},
	)

	for stack != nil {
		for stack.Val.Left != nil && !isProcessed[stack.Val.Left] {
			stack = push(
				stack,
				&StackNode{
					Val: stack.Val.Left,
				},
			)
		}
		top := stack.Val
		if k <= 1 {
			return top.Val
		}
		k--
		stack = pop(stack)
		isProcessed[top] = true
		if top.Right != nil {
			stack = push(
				stack,
				&StackNode{
					Val: top.Right,
				},
			)
		}
	}

	return 0
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type StackNode struct {
	Val  *TreeNode
	Next *StackNode
}

func push(stack *StackNode, node *StackNode) *StackNode {
	node.Next = stack
	return node
}

func pop(stack *StackNode) *StackNode {
	node := stack
	stack = stack.Next
	node.Next = nil

	return stack
}
