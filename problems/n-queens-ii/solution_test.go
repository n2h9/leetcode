package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	n int

	expected int
}

var data []testData = []testData{
	testData{
		n: 4,
		expected: 2,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, totalNQueens(item.n))
	}
}
