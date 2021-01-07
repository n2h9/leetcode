package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	root *TreeNode

	expected int
}

var data []testData = []testData{
	testData{
		root:        &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:	7,
				Left: &TreeNode{
					Val:	7,
				},
				Right: &TreeNode{
					Val:	-8,
				},
			},
			Right: &TreeNode{
				Val:	0,
			},
		},
		expected: 2,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, maxLevelSum(item.root))
	}
}
