package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	A int

	expected int
}

var data []testData = []testData{
	testData{
		A:        2,
		expected: 3,
	},
	testData{
		A:        1,
		expected: 2,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, f(item.A))
	}
}
