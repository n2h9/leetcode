package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	A []int

	expected int
}

var data []testData = []testData{
	testData{
		A: []int{1, 2, 3, 4},

		expected: 3,
	},
	testData{
		A: []int{1, 2, 8, 15},

		expected: 0,
	},
	testData{
		A: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},

		expected: 36,
	},
	testData{
		A: []int{1, 2, 8, 14, 15},

		expected: 1,
	},
	testData{
		A: []int{1, 2},

		expected: 0,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, numberOfArithmeticSlices(item.A))
	}
}
