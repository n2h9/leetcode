package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	A := &TreeNode{Val: 3}
	A.Left = &TreeNode{Val: 5}
	A.Right = &TreeNode{Val: 1}
	A.Left.Left = &TreeNode{Val: 6}
	A.Left.Right = &TreeNode{Val: 2}
	A.Left.Right.Left = &TreeNode{Val: 7}
	A.Left.Right.Right = &TreeNode{Val: 4}
	A.Right.Left = &TreeNode{Val: 0}
	A.Right.Right = &TreeNode{Val: 8}
	expected := 2
	assert.Equal(t, expected, subtreeWithAllDeepest(A).Val)
}

func TestSol1(t *testing.T) {
	A := &TreeNode{Val: 3}
	A.Left = &TreeNode{Val: 5}
	A.Right = &TreeNode{Val: 1}
	A.Left.Left = &TreeNode{Val: 6}
	A.Left.Right = &TreeNode{Val: 2}
	A.Left.Right.Left = &TreeNode{Val: 7}
	A.Left.Right.Right = &TreeNode{Val: 4}
	A.Right.Left = &TreeNode{Val: 0}
	A.Right.Right = &TreeNode{Val: 8}
	A.Right.Right.Left = &TreeNode{Val: 9}
	A.Right.Right.Right = &TreeNode{Val: 10}
	expected := 3
	assert.Equal(t, expected, subtreeWithAllDeepest(A).Val)
}

func TestSol2(t *testing.T) {
	A := &TreeNode{Val: 3}
	A.Left = &TreeNode{Val: 5}
	A.Right = &TreeNode{Val: 1}
	A.Left.Left = &TreeNode{Val: 6}
	A.Left.Left.Left = &TreeNode{Val: 9}
	A.Left.Right = &TreeNode{Val: 2}
	A.Left.Right.Left = &TreeNode{Val: 7}
	A.Left.Right.Right = &TreeNode{Val: 4}
	A.Right.Left = &TreeNode{Val: 0}
	A.Right.Right = &TreeNode{Val: 8}
	expected := 5
	assert.Equal(t, expected, subtreeWithAllDeepest(A).Val)
}

func TestSol3(t *testing.T) {
	A := &TreeNode{Val: 3}
	A.Left = &TreeNode{Val: 5}
	A.Right = &TreeNode{Val: 1}
	A.Left.Left = &TreeNode{Val: 6}
	A.Left.Right = &TreeNode{Val: 2}
	A.Left.Right.Left = &TreeNode{Val: 7}
	A.Right.Left = &TreeNode{Val: 0}
	A.Right.Right = &TreeNode{Val: 8}
	expected := 7
	assert.Equal(t, expected, subtreeWithAllDeepest(A).Val)
}
