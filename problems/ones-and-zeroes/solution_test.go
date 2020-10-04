package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	strs []string
	m    int
	n    int

	expected int
}

var data []testData = []testData{
	testData{
		strs:     []string{"10", "0001", "111001", "1", "0"},
		m:        5,
		n:        3,
		expected: 4,
	},
	testData{
		strs:     []string{"10", "1", "0"},
		m:        1,
		n:        1,
		expected: 2,
	},
	testData{
		strs:     []string{"10", "1", "0", "1111", "00", "11", "1100"},
		m:        5,
		n:        3,
		expected: 4,
	},
	testData{
		strs:     []string{"0", "00", "1"},
		m:        1,
		n:        0,
		expected: 1,
	},
	testData{
		strs:     []string{"0", "00", "1"},
		m:        0,
		n:        0,
		expected: 0,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, findMaxForm(item.strs, item.m, item.n))
	}
}
