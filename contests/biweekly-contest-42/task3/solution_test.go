package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	binary string

	expected string
}

var data []testData = []testData{
	testData{
		binary: "000110",
		expected: "111011",
	},
	testData{
		binary: "01",
		expected: "01",
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, maximumBinaryString(item.binary))
	}
}
