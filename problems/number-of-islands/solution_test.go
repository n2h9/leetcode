package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	grid [][]byte

	expected int
}

var data []testData = []testData{
	testData{
		grid: [][]byte{
			[]byte{'1', '1', '1', '1', '0'},
			[]byte{'1', '1', '0', '1', '0'},
			[]byte{'1', '1', '0', '0', '0'},
			[]byte{'0', '0', '0', '0', '0'},
		},
		expected: 1,
	},
	testData{
		grid: [][]byte{
			[]byte{'1', '1', '0', '0', '0'},
			[]byte{'1', '1', '0', '0', '0'},
			[]byte{'0', '0', '1', '0', '0'},
			[]byte{'0', '0', '0', '1', '1'},
		},
		expected: 3,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, numIslands(item.grid))
	}
}
