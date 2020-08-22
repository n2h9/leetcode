package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	root := &TreeNode{}
	root.Left = &TreeNode{}
	root.Right = &TreeNode{}
	root.Left.Right = &TreeNode{}
	distance := 3
	expected := 1

	assert.Equal(t, expected, countPairs(root, distance))
}

func TestSol1(t *testing.T) {
	root := &TreeNode{}
	root.Left = &TreeNode{}
	root.Right = &TreeNode{}
	root.Left.Right = &TreeNode{}
	root.Left.Left = &TreeNode{}
	root.Right.Right = &TreeNode{}
	root.Right.Left = &TreeNode{}
	distance := 3
	expected := 2

	assert.Equal(t, expected, countPairs(root, distance))
}

func TestSol2(t *testing.T) {
	root := &TreeNode{}
	distance := 1
	expected := 0

	assert.Equal(t, expected, countPairs(root, distance))
}

func TestSol3(t *testing.T) {
	distance := 0
	expected := 0

	assert.Equal(t, expected, countPairs(nil, distance))
}

func TestSol4(t *testing.T) {
	root := &TreeNode{}
	root.Left = &TreeNode{}
	root.Right = &TreeNode{}
	root.Left.Right = &TreeNode{}
	root.Left.Left = &TreeNode{}
	root.Right.Right = &TreeNode{}
	root.Right.Left = &TreeNode{}
	root.Right.Right.Left = &TreeNode{}
	root.Right.Right.Right = &TreeNode{}
	root.Right.Right.Right.Right = &TreeNode{}
	root.Right.Right.Right.Right.Right = &TreeNode{}
	root.Right.Right.Right.Right.Left = &TreeNode{}
	distance := 8
	expected := 15

	assert.Equal(t, expected, countPairs(root, distance))
}
