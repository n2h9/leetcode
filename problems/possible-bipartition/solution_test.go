package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	N        int
	dislikes [][]int

	expected bool
}

var data []testData = []testData{
	testData{
		dislikes: [][]int{
			[]int{1, 2}, []int{1, 3}, []int{2, 4},
		},
		N: 4,

		expected: true,
	},
	testData{
		dislikes: [][]int{
			[]int{1, 2}, []int{1, 3}, []int{2, 3},
		},
		N: 3,

		expected: false,
	},
	testData{
		dislikes: [][]int{
			[]int{1, 2}, []int{2, 3}, []int{3, 4}, []int{4, 5}, []int{1, 5},
		},
		N: 5,

		expected: false,
	},
	testData{
		dislikes: [][]int{
			[]int{1, 2}, []int{2, 3}, []int{2, 4}, []int{4, 5}, []int{1, 5},
		},
		N: 5,

		expected: true,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, possibleBipartition(item.N, item.dislikes))
	}
}
