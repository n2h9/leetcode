package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	nums []int

	expected [][]int
}

var data []testData = []testData{
	testData{
		nums: []int{1, 2, 3},

		expected: [][]int{
			[]int{1, 2, 3},
			[]int{1, 3, 2},
			[]int{2, 1, 3},
			[]int{2, 3, 1},
			[]int{3, 1, 2},
			[]int{3, 2, 1},
		},
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, permute(item.nums))
	}
}
