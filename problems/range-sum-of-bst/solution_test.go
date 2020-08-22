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
	relations := []int{10, 5, 15, 3, 7, -1, 18}
	root := &TreeNode{
		Val: relations[0],
	}
	TreeFromArray(root, 0, relations)
	L := 7
	R := 15

	expected := 32

	assert.Equal(t, expected, rangeSumBST(root, L, R))
}

func TestSol1(t *testing.T) {
	relations := []int{10, 5, 15, 3, 7, 13, 18, 1, -1, 6}
	root := &TreeNode{
		Val: relations[0],
	}
	TreeFromArray(root, 0, relations)
	L := 6
	R := 10

	expected := 23

	assert.Equal(t, expected, rangeSumBST(root, L, R))
}

func TestSol2(t *testing.T) {
	relations := []int{10, 5, 15, 3, 7, 13, 18, 1, -1, 6}
	root := &TreeNode{
		Val: relations[0],
	}
	TreeFromArray(root, 0, relations)
	L := 0
	R := 100

	expected := 78

	assert.Equal(t, expected, rangeSumBST(root, L, R))
}

func TestSol3(t *testing.T) {
	relations := []int{10, 5, 15, 3, 7, 13, 18, 1, -1, 6}
	root := &TreeNode{
		Val: relations[0],
	}
	TreeFromArray(root, 0, relations)
	L := 11
	R := 17

	expected := 28

	assert.Equal(t, expected, rangeSumBST(root, L, R))
}

func TestSol4(t *testing.T) {
	relations := []int{10, 5, 15, 3, 7, 13, 18, 1, -1, 6}
	root := &TreeNode{
		Val: relations[0],
	}
	TreeFromArray(root, 0, relations)
	L := 4
	R := 9

	expected := 18

	assert.Equal(t, expected, rangeSumBST(root, L, R))
}

func TestSol5(t *testing.T) {
	relations := []int{10, 5, 15, 3, 7, 13, 18, 1, -1, 6}
	root := &TreeNode{
		Val: relations[0],
	}
	TreeFromArray(root, 0, relations)
	L := 8
	R := 9

	expected := 0

	assert.Equal(t, expected, rangeSumBST(root, L, R))
}

func TestSol6(t *testing.T) {
	relations := []int{10, 5, 15, 3, 7, 13, 18, 1, -1, 6}
	root := &TreeNode{
		Val: relations[0],
	}
	TreeFromArray(root, 0, relations)
	L := 7
	R := 7

	expected := 7

	assert.Equal(t, expected, rangeSumBST(root, L, R))
}

func TestSol7(t *testing.T) {
	relations := []int{18, 9, 27, 6, 15, 24, 30, 3, -1, 12, -1, 21}
	root := &TreeNode{
		Val: relations[0],
	}
	TreeFromArray(root, 0, relations)
	L := 18
	R := 24

	expected := 63

	assert.Equal(t, expected, rangeSumBST(root, L, R))
}
