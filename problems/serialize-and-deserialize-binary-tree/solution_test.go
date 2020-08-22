package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	codec := Constructor()
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
	}
	s := codec.serialize(root)
	rootD := codec.deserialize(s)
	assert.Equal(t, root.Left.Val, rootD.Left.Val)
	assert.Equal(t, root.Right.Right.Val, rootD.Right.Right.Val)
}
