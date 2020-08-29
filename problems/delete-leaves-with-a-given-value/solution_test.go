package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 2}

	root = removeLeafNodes(root, 2)

	assert.Equal(t, 1, root.Val)
	assert.Nil(t, root.Left)
}
