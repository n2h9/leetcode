package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	A []int
	K int

	expected int
}

var data []testData = []testData{
	testData{
		A:        []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
		K:        2,
		expected: 6,
	},
	testData{
		A:        []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
		K:        3,
		expected: 10,
	},
	testData{
		A:        []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
		K:        0,
		expected: 4,
	},
	testData{
		A:        []int{0, 0, 0, 0},
		K:        0,
		expected: 0,
	},
	testData{
		A:        []int{1},
		K:        0,
		expected: 1,
	},
	testData{
		A:        []int{},
		K:        0,
		expected: 0,
	},
	testData{
		A:        []int{1, 1, 1, 1, 1, 1, 1, 1},
		K:        0,
		expected: 8,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, longestOnes(item.A, item.K))
	}
}
