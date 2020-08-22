package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	root := &TreeNode{Val: 1}
	root2 := &TreeNode{Val: 2}
	root4 := &TreeNode{Val: 4}

	root.Left = root2
	root.Right = &TreeNode{Val: 3}
	root2.Left = root4

	root.Right.Left = root2
	root.Right.Right = root4

	expected := []int{4, 2}

	result := []int{}
	list := findDuplicateSubtrees(root)
	for _, l := range list {
		result = append(result, l.Val)
	}

	assert.Equal(t, expected, result)
}

func TestSol1(t *testing.T) {
	root := &TreeNode{Val: 1}
	expected := []*TreeNode{}

	assert.Equal(t, expected, findDuplicateSubtrees(root))
}
