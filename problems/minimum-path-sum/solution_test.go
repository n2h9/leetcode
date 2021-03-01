package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	type testData struct {
		grid [][]int

		expected int
	}

	var data []testData = []testData{
		testData{
			grid: [][]int{
				[]int{1, 3, 1},
				[]int{1, 5, 1},
				[]int{4, 2, 1},
			},
			expected: 7,
		},
		testData{
			grid: [][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
			},
			expected: 12,
		},
		testData{
			grid: [][]int{
				[]int{1, 2, 3},
			},
			expected: 6,
		},
		testData{
			grid: [][]int{
				[]int{1},
				[]int{2},
				[]int{3},
			},
			expected: 6,
		},
		testData{
			grid: [][]int{
				[]int{33},
			},
			expected: 33,
		},
	}

	for _, item := range data {
		assert.Equal(t, item.expected, minPathSum(item.grid))
	}
}
