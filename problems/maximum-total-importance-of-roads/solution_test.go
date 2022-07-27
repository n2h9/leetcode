package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	n     int
	roads [][]int

	expected int64
}

var data []testData = []testData{
	testData{
		n:     5,
		roads: [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 2}, {1, 3}, {2, 4}},

		expected: 43,
	},

	testData{
		n:     5,
		roads: [][]int{{0, 3}, {2, 4}, {1, 3}},

		expected: 20,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, maximumImportance(item.n, item.roads))
	}
}
