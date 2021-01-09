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
		n:        4,
		expected: 10,
	},
	testData{
		n:			10,
		expected: 37,
	},
	testData{
		n:			20,
		expected: 96,
	},
	testData{
		n:			1,
		expected: 1,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, totalMoney(item.n))
	}
}
