package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	n int

	expected [][]string
}

var data []testData = []testData{
	testData{
		n: 4,
		expected: [][]string{
			[]string{
				".Q..",
				"...Q",
				"Q...",
				"..Q.",
			},
			[]string{
				"..Q.",
				"Q...",
				"...Q",
				".Q..",
			},
		},
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, solveNQueens(item.n))
	}
}
