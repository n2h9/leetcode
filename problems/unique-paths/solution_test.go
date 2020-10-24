package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	m int
	n int

	expected int
}

var data []testData = []testData{
	testData{
		m:        3,
		n:        2,
		expected: 3,
	},
	testData{
		m:        3,
		n:        7,
		expected: 28,
	},
	testData{
		m:        3,
		n:        3,
		expected: 6,
	},
	testData{
		m:        7,
		n:        3,
		expected: 28,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, uniquePaths(item.m, item.n))
	}
}
