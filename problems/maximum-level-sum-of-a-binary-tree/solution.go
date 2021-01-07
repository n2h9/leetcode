package solution

type TreeNode struct {
	Val int
	Left *TreeNode
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
func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxSum, maxSumLevelNo := ^(1 << 32)+1, 0
	q := newQueue()

	q.push(newNode(root))
	level := 0
	for !q.isEmpty() {
		level++
		qNextLevel := newQueue()
		sum := 0
		for !q.isEmpty() {
			n := q.top()
			q.pop()
			sum += n.val.Val
			if n.val.Left != nil {
				qNextLevel.push(newNode(n.val.Left))
			}
			if n.val.Right != nil {
				qNextLevel.push(newNode(n.val.Right))
			}		
		}
		q = qNextLevel
		if sum > maxSum {
			maxSum = sum
			maxSumLevelNo = level
		}
	}
	return maxSumLevelNo
}

type node struct {
	val *TreeNode
	next * node
}

func newNode(val *TreeNode) *node {
	return &node{
		val: val,
	}
}

type queue struct {
	head *node
	tail *node
}

func newQueue() *queue {
	return &queue{}
}

func (q *queue) isEmpty() bool {
	return q.head == nil
}

func (q *queue) top() *node {
	return q.head
}

func (q *queue) pop() {
	n := q.head
	q.head = q.head.next
	n.next = nil

	if q.head == nil {
		q.tail = nil
	}
}

func (q *queue) push(n *node) {
	if q.tail == nil {
		q.tail = n
		q.head = n
		return
	}

	q.tail.next = n
	q.tail = q.tail.next
}