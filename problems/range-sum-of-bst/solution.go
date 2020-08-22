package solution

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}

	stack := push(
		nil,
		&StackNode{
			Val: root,
		},
	)

	// 1) search for L
	for {
		if L == stack.Val.Val {
			// L is current node
			break
		}
		var node *TreeNode
		if L < stack.Val.Val {
			node = stack.Val.Left
		} else {
			node = stack.Val.Right
		}

		if node == nil {
			// stop search, found nearest to L value
			break
		}
		stack = push(
			stack,
			&StackNode{
				Val: node,
			},
		)
	}
	// --

	// 2) sum and search for R
	sum := 0
	isProcessed := make(map[*TreeNode]bool)
	for stack != nil {
		top := stack
		stack = pop(stack)
		isProcessed[top.Val] = true
		if top.Val.Val >= L && top.Val.Val <= R {
			sum += top.Val.Val
		}
		if top.Val.Val < R && top.Val.Right != nil && !isProcessed[top.Val.Right] {
			stack = push(
				stack,
				&StackNode{
					Val: top.Val.Right,
				},
			)
		}
		if top.Val.Val > L && top.Val.Left != nil && !isProcessed[top.Val.Left] {
			stack = push(
				stack,
				&StackNode{
					Val: top.Val.Left,
				},
			)
		}
	}
	// --

	return sum
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
