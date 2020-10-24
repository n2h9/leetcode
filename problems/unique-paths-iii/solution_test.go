package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	grid [][]int

	expected int
}

var data []testData = []testData{
	testData{
		grid: [][]int{
			[]int{1, 0, 0, 0},
			[]int{0, 0, 0, 0},
			[]int{0, 0, 2, -1},
		},
		expected: 2,
	},
	testData{
		grid: [][]int{
			[]int{1, 0, 0, 0},
			[]int{0, 0, 0, 0},
			[]int{0, 0, 0, 2},
		},
		expected: 4,
	},
	testData{
		grid:     [][]int{[]int{0, 1}, []int{2, 0}},
		expected: 0,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, uniquePathsIII(item.grid))
	}
}
