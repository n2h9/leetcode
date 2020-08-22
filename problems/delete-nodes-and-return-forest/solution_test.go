package solution

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	toDelete := []int{3, 5}
	expected := []int{1, 6, 7}

	forest := delNodes(root, toDelete)
	res := make([]int, 0)
	for _, tree := range forest {
		res = append(res, tree.Val)
	}

	assert.Equal(t, expected, res)
}

func TestSol1(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	toDelete := []int{}
	expected := []int{1}

	forest := delNodes(root, toDelete)
	res := make([]int, 0)
	for _, tree := range forest {
		res = append(res, tree.Val)
	}

	assert.Equal(t, expected, res)
}

func TestSol2(t *testing.T) {
	toDelete := []int{3, 5}
	expected := []int{}

	forest := delNodes(nil, toDelete)
	res := make([]int, 0)
	for _, tree := range forest {
		res = append(res, tree.Val)
	}

	assert.Equal(t, expected, res)
}

func TestSol3(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Left.Left.Left = &TreeNode{Val: 8}
	root.Left.Left.Right = &TreeNode{Val: 9}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	toDelete := []int{3, 2, 4}
	expected := []int{1, 5, 6, 7, 8, 9}

	forest := delNodes(root, toDelete)
	res := make([]int, 0)
	for _, tree := range forest {
		res = append(res, tree.Val)
	}
	sort.Ints(res)

	assert.Equal(t, expected, res)
}

func TestSol4(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Left.Left.Left = &TreeNode{Val: 8}
	root.Left.Left.Right = &TreeNode{Val: 9}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	toDelete := []int{1, 2, 3, 4, 5}
	expected := []int{6, 7, 8, 9}

	forest := delNodes(root, toDelete)
	res := make([]int, 0)
	for _, tree := range forest {
		res = append(res, tree.Val)
	}
	sort.Ints(res)

	assert.Equal(t, expected, res)
}
