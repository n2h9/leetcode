package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	nums []int

	expected int
}

var data []testData = []testData{
	testData{
		nums:     []int{3, 4, 2},
		expected: 6,
	},
	testData{
		nums:     []int{2, 2, 3, 3, 3, 4},
		expected: 9,
	},
	testData{
		nums:     []int{2, 2, 3, 3, 3, 4, 6, 6},
		expected: 21,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, deleteAndEarn(item.nums))
	}
}
