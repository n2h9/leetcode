package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	grid     [][]int
	moveCost [][]int

	expected int
}

var data []testData = []testData{
	testData{
		grid: [][]int{
			{5, 3}, {4, 0}, {2, 1},
		},
		moveCost: [][]int{
			{9, 8}, {1, 5}, {10, 12}, {18, 6}, {2, 4}, {14, 3},
		},
		expected: 17,
	},
	testData{
		grid: [][]int{
			{5, 1, 2}, {4, 0, 3},
		},
		moveCost: [][]int{
			{12, 10, 15}, {20, 23, 8}, {21, 7, 1}, {8, 1, 13}, {9, 10, 25}, {5, 3, 2},
		},
		expected: 6,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, minPathCost(item.grid, item.moveCost))
	}
}
