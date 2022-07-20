package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	parents []int

	expected int
}

var data []testData = []testData{
	testData{
		parents:  []int{-1, 2, 0, 2, 0},
		expected: 3,
	},
	testData{
		parents:  []int{-1, 12, 16, 0, 0, 7, 17, 14, 17, 16, 9, 13, 14, 4, 4, 13, 3, 3, 8, 11},
		expected: 1,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, countHighestScoreNodes(item.parents))
	}
}
