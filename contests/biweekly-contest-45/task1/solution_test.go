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
			nums:     []int{1, 2, 3, 2},
			expected: 4,
		},
		testData{
			nums:     []int{1, 1, 1, 1},
			expected: 0,
		},
		testData{
			nums:     []int{1, 2, 3, 4, 5},
			expected: 15,
		},
	}

	for _, item := range data {
		assert.Equal(t, item.expected, sumOfUnique(item.nums))
	}
}
