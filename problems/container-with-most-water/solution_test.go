package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	height []int

	expected int
}

var data []testData = []testData{
	testData{
		height:        []int{1,8,6,2,5,4,8,3,7},
		expected: 49,
	},
	testData{
		height:        []int{1,1},
		expected: 1,
	},
	testData{
		height:        []int{4,3,2,1,4},
		expected: 16,
	},
	testData{
		height:        []int{1,2,1},
		expected: 2,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, maxArea(item.height))
	}
}
