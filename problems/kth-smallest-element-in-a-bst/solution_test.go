package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const NO_CHILD_VALUE = -1

// [5,3,6,2,4,-1,-1,1] -1 means no child
func TreeFromArray(root *TreeNode, index int, relations []int) {
	li := 2*index + 1
	if li < len(relations) && relations[li] != NO_CHILD_VALUE {
		root.Left = &TreeNode{
			Val: relations[li],
		}
		TreeFromArray(root.Left, li, relations)
	}
	ri := 2*index + 2
	if ri < len(relations) && relations[ri] != NO_CHILD_VALUE {
		root.Right = &TreeNode{
			Val: relations[ri],
		}
		TreeFromArray(root.Right, ri, relations)
	}
}

func TestSol0(t *testing.T) {
	relations := []int{3, 1, 4, -1, 2}
	root := &TreeNode{
		Val: relations[0],
	}
	TreeFromArray(root, 0, relations)
	k := 1

	expected := 1

	assert.Equal(t, expected, kthSmallest(root, k))
}

func TestSol1(t *testing.T) {
	relations := []int{5, 3, 6, 2, 4, -1, -1, 1}
	root := &TreeNode{
		Val: relations[0],
	}
	TreeFromArray(root, 0, relations)
	k := 3

	expected := 3

	assert.Equal(t, expected, kthSmallest(root, k))
}

func TestSol2(t *testing.T) {
	relations := []int{10, 5, 20, 2, 7, 15, 25, 1, 3}
	root := &TreeNode{
		Val: relations[0],
	}
	TreeFromArray(root, 0, relations)
	k := 9

	expected := 25

	assert.Equal(t, expected, kthSmallest(root, k))
}

func TestSol3(t *testing.T) {
	relations := []int{10, 5, 20, 2, 7, 15, 25, 1, 3, -1, -1, -1, 17}
	root := &TreeNode{
		Val: relations[0],
	}
	TreeFromArray(root, 0, relations)
	k := 8

	expected := 17

	assert.Equal(t, expected, kthSmallest(root, k))
}
