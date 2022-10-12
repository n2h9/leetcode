package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	words []string

	expected []int
}

var data []testData = []testData{
	{
		words:    []string{"abc", "ab", "bc", "b"},
		expected: []int{5, 4, 3, 2},
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, sumPrefixScores(item.words))
	}
}
