package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	times [][]int
	N     int
	K     int

	expected int
}

var data []testData = []testData{
	testData{
		times: [][]int{
			[]int{2, 1, 1}, []int{2, 3, 1}, []int{3, 4, 1},
		},
		N: 4,
		K: 2,

		expected: 2,
	},
	testData{
		times: [][]int{
			[]int{2, 1, 1}, []int{2, 3, 1}, []int{3, 4, 1}, []int{5, 6, 1},
		},
		N: 6,
		K: 2,

		expected: -1,
	},
	testData{
		times: [][]int{
			[]int{2, 1, 8}, []int{2, 3, 1}, []int{3, 4, 1}, []int{4, 1, 1},
		},
		N: 4,
		K: 2,

		expected: 3,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, networkDelayTime(item.times, item.N, item.K))
	}
}
