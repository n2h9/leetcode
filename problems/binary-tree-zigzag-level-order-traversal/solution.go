package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	yellow, blue := newStack(), newStack()
	if root != nil {
		yellow.push(newNode(root))
	}
	order := make([][]int, 0)
	for !yellow.isEmpty() {
		levelOrder := make([]int, 0)
		for !yellow.isEmpty() {
			treeNode := yellow.top().val
			yellow.pop()
			levelOrder = append(levelOrder, treeNode.Val)
			for _, tn := range []*TreeNode{treeNode.Left, treeNode.Right} {
				if tn != nil {
					blue.push(newNode(tn))
				}
			}
		}
		if len(levelOrder) > 0 {
			order = append(order, levelOrder)
		}

		levelOrder = make([]int, 0)
		for !blue.isEmpty() {
			treeNode := blue.top().val
			blue.pop()
			levelOrder = append(levelOrder, treeNode.Val)
			// (!) important the order of child addition to stack is changed (right child first)
			for _, tn := range []*TreeNode{treeNode.Right, treeNode.Left} {
				if tn != nil {
					yellow.push(newNode(tn))
				}
			}
		}
		if len(levelOrder) > 0 {
			order = append(order, levelOrder)
		}
	}
	return order
}

type node struct {
	val  *TreeNode
	next *node
}

func newNode(val *TreeNode) *node {
	return &node{val: val}
}

type stack struct {
	head *node
}

func newStack() *stack {
	return &stack{}
}

func (s *stack) isEmpty() bool {
	return s.head == nil
}

func (s *stack) top() *node {
	return s.head
}

func (s *stack) pop() {
	head := s.head
	s.head = s.head.next
	head.next = nil
}

func (s *stack) push(n *node) {
	n.next = s.head
	s.head = n
}
