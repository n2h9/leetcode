package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	digits string

	expected []string
}

var data []testData = []testData{
	testData{
		digits:        "23",
		expected: []string{"ad","ae","af","bd","be","bf","cd","ce","cf"},
	},
	testData{
		digits: "",
		expected: []string{},
	},
	testData{
		digits: "2",
		expected: []string{"a", "b", "c"},
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, letterCombinations(item.digits))
	}
}
