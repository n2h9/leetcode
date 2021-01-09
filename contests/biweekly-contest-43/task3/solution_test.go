package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	n int

	expected []int
}

var data []testData = []testData{
	testData{
		n:        3,
		expected: []int{3,1,2,3,2},
	},
	testData{
		n:        3,
		expected: []int{5,3,1,4,3,5,2,4,2},
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, constructDistancedSequence(item.n))
	}
}
