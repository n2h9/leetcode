package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	matrix   [][]int
	expected int
}

var data []testData = []testData{
	testData{
		matrix: [][]int{
			[]int{0, 1, 1, 1},
			[]int{1, 1, 1, 1},
			[]int{0, 1, 1, 1},
		},

		expected: 15,
	},
	testData{
		matrix: [][]int{
			[]int{1, 0, 1},
			[]int{1, 1, 0},
			[]int{1, 1, 0},
		},

		expected: 7,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, countSquares(item.matrix))
	}
}
