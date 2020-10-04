package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	triangle [][]int

	expected int
}

var data []testData = []testData{
	testData{
		triangle: [][]int{[]int{2}, []int{3, 4}, []int{6, 5, 7}, []int{4, 1, 8, 3}},
		expected: 11,
	},
	testData{
		triangle: [][]int{[]int{7}, []int{3, 4}},
		expected: 10,
	},
	testData{
		triangle: [][]int{[]int{7}},
		expected: 7,
	},
	testData{
		triangle: [][]int{},
		expected: 0,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, minimumTotal(item.triangle))
	}
}
