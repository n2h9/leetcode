package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	n      int
	edges  [][]int
	labels string

	expected []int
}

var data []testData = []testData{
	testData{
		n:      7,
		edges:  [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}},
		labels: "abaedcd",

		expected: []int{2, 1, 1, 1, 1, 1, 1},
	},
	testData{
		n:      4,
		edges:  [][]int{{0, 1}, {1, 2}, {0, 3}},
		labels: "bbbb",

		expected: []int{4, 2, 1, 1},
	},
	testData{
		n:      4,
		edges:  [][]int{{0, 2}, {0, 3}, {1, 2}},
		labels: "aeed",

		expected: []int{1, 1, 2, 1},
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, countSubTrees(item.n, item.edges, item.labels))
	}
}
