package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	points [][]int
	k int

	expected [][]int
}

var data []testData = []testData{
	testData{
		points:	[][]int{[]int{1,3},[]int{-2,2}},
		k: 1,
		expected: [][]int{[]int{-2,2}},
	},
	testData{
		points:	[][]int{[]int{3,3},[]int{5,-1},[]int{-2,4}},
		k: 2,
		expected: [][]int{[]int{3,3}, []int{-2,4}},
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, kClosest(item.points, item.k))
	}
}
