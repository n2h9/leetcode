package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	graph [][]int

	expected bool
}

var data []testData = []testData{
	testData{
		graph:    [][]int{[]int{1, 3}, []int{0, 2}, []int{1, 3}, []int{0, 2}},
		expected: true,
	},
	testData{
		graph:    [][]int{[]int{1, 2, 3}, []int{0, 2}, []int{0, 1, 3}, []int{0, 2}},
		expected: false,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, isBipartite(item.graph))
	}
}
