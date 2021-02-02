package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	time []int

	expected int
}

var data []testData = []testData{
	testData{
		time:     []int{30, 20, 150, 100, 40},
		expected: 3,
	},
	testData{
		time:     []int{60, 60, 60},
		expected: 3,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, numPairsDivisibleBy60(item.time))
	}
}
