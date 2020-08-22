package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func NodeToSlice(node *ListNode) []int {
	v := []int{}
	for node != nil {
		v = append(v, node.Val)
		node = node.Next
	}
	return v
}

func TestSol0(t *testing.T) {
	node := &ListNode{
		Val: 4,
	}
	node.Next = &ListNode{Val: 5}
	node.Next.Next = &ListNode{Val: 1}
	node.Next.Next.Next = &ListNode{Val: 9}

	expected := []int{4, 1, 9}
	deleteNode(node.Next)
	assert.Equal(t, expected, NodeToSlice(node))
}

func TestSol1(t *testing.T) {
	node := &ListNode{
		Val: 4,
	}
	node.Next = &ListNode{Val: 5}
	node.Next.Next = &ListNode{Val: 1}
	node.Next.Next.Next = &ListNode{Val: 9}

	expected := []int{4, 5, 9}
	deleteNode(node.Next.Next)
	assert.Equal(t, expected, NodeToSlice(node))
}

func TestSol2(t *testing.T) {
	node := &ListNode{
		Val: 4,
	}
	node.Next = &ListNode{Val: 5}

	expected := []int{5}
	deleteNode(node)
	assert.Equal(t, expected, NodeToSlice(node))
}
