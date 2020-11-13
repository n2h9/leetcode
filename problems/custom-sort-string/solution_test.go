package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	S string
	T string

	expected string
}

var data []testData = []testData{
	testData{
		S: "cba",
		T: "abcd",

		expected: "cbad",
	},
	testData{
		S: "cba",
		T: "abcdefggfedcba",

		expected: "ccbbaadefggfed",
	},
	testData{
		S: "",
		T: "xyz",

		expected: "xyz",
	},
	testData{
		S: "abc",
		T: "",

		expected: "",
	},
	testData{
		S: "",
		T: "",

		expected: "",
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, customSortString(item.S, item.T))
	}
}
