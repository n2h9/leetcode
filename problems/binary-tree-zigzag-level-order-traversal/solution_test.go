package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}
	expected := [][]int{
		[]int{3},
		[]int{20, 9},
		[]int{15, 7},
	}

	assert.Equal(t, expected, zigzagLevelOrder(root))
}

func TestSol1(t *testing.T) {
	root := &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 20}
	root.Right.Right = &TreeNode{Val: 7}
	root.Right.Right.Right = &TreeNode{Val: 90}
	expected := [][]int{
		[]int{3},
		[]int{20},
		[]int{7},
		[]int{90},
	}

	assert.Equal(t, expected, zigzagLevelOrder(root))
}

func TestSol2(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Right = &TreeNode{Val: 3}
	expected := [][]int{
		[]int{1},
		[]int{2},
		[]int{3},
	}

	assert.Equal(t, expected, zigzagLevelOrder(root))
}
