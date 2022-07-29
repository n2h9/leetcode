package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	n     int
	edges [][]int

	expected int64
}

var data []testData = []testData{
	testData{
		n:        3,
		edges:    [][]int{{0, 1}, {0, 2}, {1, 2}},
		expected: 0,
	},
	testData{
		n:        7,
		edges:    [][]int{{0, 2}, {0, 5}, {2, 4}, {1, 6}, {5, 4}},
		expected: 14,
	},
	testData{
		n:        11,
		edges:    [][]int{{0, 2}, {0, 5}, {2, 4}, {1, 6}, {5, 4}, {7, 8}, {7, 9}, {8, 10}, {9, 10}},
		expected: 42,
	},
	testData{
		n:        100000,
		edges:    [][]int{},
		expected: 4999950000,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, countPairs(item.n, item.edges))
	}
}
