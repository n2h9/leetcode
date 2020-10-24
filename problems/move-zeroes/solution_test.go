package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	nums []int

	expected []int
}

var data []testData = []testData{
	testData{
		nums:     []int{0, 1, 0, 3, 12},
		expected: []int{1, 3, 12, 0, 0},
	},
	testData{
		nums:     []int{},
		expected: []int{},
	},
	testData{
		nums:     []int{0, 0, 0, 0, 0},
		expected: []int{0, 0, 0, 0, 0},
	},
	testData{
		nums:     []int{1, 2, 3, 4, 1},
		expected: []int{1, 2, 3, 4, 1},
	},
	testData{
		nums:     []int{0},
		expected: []int{0},
	},
	testData{
		nums:     []int{1},
		expected: []int{1},
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		moveZeroes(item.nums)
		assert.Equal(t, item.expected, item.nums)
	}
}
