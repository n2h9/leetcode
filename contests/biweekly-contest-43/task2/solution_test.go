package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	s string
	x int
	y int

	expected int
}

var data []testData = []testData{
	testData{
		s:        "cdbcbbaaabab",
		x: 4,
		y: 5,
		expected: 19,
	},
	testData{
		s: "aabbaaxybbaabb",
		x: 5,
		y: 4,
		expected: 20,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, maximumGain(item.s, item.x, item.y))
	}
}
