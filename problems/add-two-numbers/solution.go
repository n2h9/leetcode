package solution

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var current *ListNode

	overflow := 0
	for l1 != nil || l2 != nil {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		sum := v1 + v2 + overflow
		if sum > 9 {
			sum %= 10
			overflow = 1
		} else {
			overflow = 0
		}

		node := &ListNode{
			Val:  sum,
			Next: nil,
		}

		if current == nil {
			head = node
			current = node
			continue
		}

		current.Next = node
		current = node
	}

	if overflow != 0 {
		current.Next = &ListNode{
			Val:  overflow,
			Next: nil,
		}
	}

	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
