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
		n:        3,
		expected: 3,
	},
	testData{
		n:        1,
		expected: 0,
	},
	testData{
		n:        2,
		expected: 2,
	},
	testData{
		n:        4,
		expected: 4,
	},
	testData{
		n:        5,
		expected: 5,
	},
	testData{
		n:        6,
		expected: 5,
	},
	testData{
		n:        7,
		expected: 7,
	},
	testData{
		n:        8,
		expected: 6,
	},
	testData{
		n:        9,
		expected: 6,
	},
	testData{
		n:        15,
		expected: 8,
	},
	testData{
		n:        45,
		expected: 11,
	},
	testData{
		n:        1000,
		expected: 21,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, minSteps(item.n))
	}
}
