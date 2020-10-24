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
			[]int{0, 0, 0},
			[]int{0, 1, 0},
			[]int{0, 0, 0},
		},
		expected: 2,
	},
	testData{
		grid: [][]int{
			[]int{1, 0, 0},
			[]int{0, 1, 0},
			[]int{0, 0, 0},
		},
		expected: 0,
	},
	testData{
		grid: [][]int{
			[]int{0, 0, 1},
			[]int{0, 1, 0},
			[]int{0, 0, 0},
		},
		expected: 1,
	},
	testData{
		grid: [][]int{
			[]int{0, 0, 0},
			[]int{0, 0, 1},
			[]int{0, 1, 0},
		},
		expected: 0,
	},
	testData{
		grid: [][]int{
			[]int{0, 1, 0, 0, 0},
			[]int{0, 1, 0, 0, 0},
			[]int{0, 0, 0, 0, 0},
			[]int{0, 1, 0, 0, 0},
		},
		expected: 3,
	},
	testData{
		grid: [][]int{
			[]int{0},
			[]int{1},
		},
		expected: 0,
	},
	testData{
		grid: [][]int{
			[]int{0, 1},
		},
		expected: 0,
	},
	testData{
		grid: [][]int{
			[]int{0, 0},
			[]int{0, 1},
		},
		expected: 0,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, uniquePathsWithObstacles(item.grid))
	}
}
