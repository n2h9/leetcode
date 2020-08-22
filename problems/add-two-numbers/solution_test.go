package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func arrayToList(a []int) *ListNode {
	if len(a) < 0 {
		return nil
	}

	head := &ListNode{
		Val:  a[0],
		Next: nil,
	}
	node := head

	for i := 1; i < len(a); i++ {
		node.Next = &ListNode{
			Val:  a[i],
			Next: nil,
		}
		node = node.Next
	}

	return head
}

func listToArray(node *ListNode) []int {
	if node == nil {
		return []int{}
	}

	a := make([]int, 0)

	for node != nil {
		a = append(a, node.Val)
		node = node.Next
	}

	return a
}

func TestSol0(t *testing.T) {
	l1 := []int{2, 4, 3}
	l2 := []int{5, 6, 4}
	expected := []int{7, 0, 8}
	assert.Equal(
		t,
		expected,
		listToArray(
			addTwoNumbers(
				arrayToList(l1),
				arrayToList(l2),
			),
		),
	)
}

func TestSol1(t *testing.T) {
	l1 := []int{9, 9, 9, 9}
	l2 := []int{1}
	expected := []int{0, 0, 0, 0, 1}
	assert.Equal(
		t,
		expected,
		listToArray(
			addTwoNumbers(
				arrayToList(l1),
				arrayToList(l2),
			),
		),
	)
}

func TestSol2(t *testing.T) {
	l1 := []int{1}
	l2 := []int{9, 9, 9, 9, 9, 9, 9, 9, 9}
	expected := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	assert.Equal(
		t,
		expected,
		listToArray(
			addTwoNumbers(
				arrayToList(l1),
				arrayToList(l2),
			),
		),
	)
}

func TestSol3(t *testing.T) {
	l1 := []int{0}
	l2 := []int{8, 6, 7, 8}
	expected := []int{8, 6, 7, 8}
	assert.Equal(
		t,
		expected,
		listToArray(
			addTwoNumbers(
				arrayToList(l1),
				arrayToList(l2),
			),
		),
	)
}

func TestSol4(t *testing.T) {
	l1 := []int{9}
	l2 := []int{9}
	expected := []int{8, 1}
	assert.Equal(
		t,
		expected,
		listToArray(
			addTwoNumbers(
				arrayToList(l1),
				arrayToList(l2),
			),
		),
	)
}

func TestSol5(t *testing.T) {
	l1 := []int{9}
	l2 := []int{1, 1}
	expected := []int{0, 2}
	assert.Equal(
		t,
		expected,
		listToArray(
			addTwoNumbers(
				arrayToList(l1),
				arrayToList(l2),
			),
		),
	)
}
