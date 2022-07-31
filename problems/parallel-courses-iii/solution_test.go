package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	n         int
	relations [][]int
	time      []int

	expected int
}

var data []testData = []testData{
	{
		n:         3,
		relations: [][]int{{1, 3}, {2, 3}},
		time:      []int{3, 2, 5},

		expected: 8,
	},
	{
		n:         5,
		relations: [][]int{{1, 5}, {2, 5}, {3, 5}, {3, 4}, {4, 5}},
		time:      []int{1, 2, 3, 4, 5},

		expected: 12,
	},
	{
		n:         6,
		relations: [][]int{{1, 5}, {2, 5}, {3, 5}, {3, 4}, {4, 5}, {4, 6}, {6, 5}},
		time:      []int{1, 2, 3, 4, 5, 30},

		expected: 42,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, minimumTime(item.n, item.relations, item.time))
	}
}
