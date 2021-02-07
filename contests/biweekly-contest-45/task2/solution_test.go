package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	type testData struct {
		nums []int

		expected int
	}

	var data []testData = []testData{
		testData{
			nums:     []int{1, -3, 2, 3, -4},
			expected: 5,
		},
		testData{
			nums:     []int{2, -5, 1, -4, 3, -2},
			expected: 8,
		},
	}

	for _, item := range data {
		assert.Equal(t, item.expected, maxAbsoluteSum(item.nums))
	}
}
