package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	type testData struct {
		s string

		expected int
	}

	var data []testData = []testData{
		testData{
			s:        "ca",
			expected: 2,
		},
		testData{
			s:        "cabaabac",
			expected: 0,
		},
		testData{
			s:        "aabccabba",
			expected: 3,
		},
	}

	for _, item := range data {
		assert.Equal(t, item.expected, minimumLength(item.s))
	}
}
