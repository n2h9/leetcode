package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	queries []string
	pattern string

	expected []bool
}

var data []testData = []testData{
	{
		queries:  []string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"},
		pattern:  "FB",
		expected: []bool{true, false, true, true, false},
	},
	{
		queries:  []string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"},
		pattern:  "FoBa",
		expected: []bool{true, false, true, false, false},
	},
	{
		queries:  []string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"},
		pattern:  "FoBaT",
		expected: []bool{false, true, false, false, false},
	},
	{
		queries:  []string{"FroameBuffera"},
		pattern:  "FoBa",
		expected: []bool{true},
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, camelMatch(item.queries, item.pattern))
	}
}
