package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	nums []int
	S    int

	expected int
}

var data []testData = []testData{
	// testData{
	// 	nums:     []int{1, 1, 1, 1, 1},
	// 	S:        3,
	// 	expected: 5,
	// },
	// testData{
	// 	nums:     []int{1, 2, 7, 9, 981},
	// 	S:        1000000000,
	// 	expected: 0,
	// },
	testData{
		nums:     []int{0, 0, 0, 0, 0, 0, 0, 0, 1},
		S:        1,
		expected: 256,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, findTargetSumWays(item.nums, item.S))
	}
}
