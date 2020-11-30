package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	s string

	expected int
}

var data []testData = []testData{
	testData{
		s:        "abcabcbb",
		expected: 3,
	},
	testData{
		s:        "bbbbb",
		expected: 1,
	},
	testData{
		s:        "",
		expected: 0,
	},
	testData{
		s: "abdadcdba",
		expected: 4,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, lengthOfLongestSubstring(item.s))
	}
}
