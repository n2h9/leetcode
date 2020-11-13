package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	s string

	expected []string
}

var data []testData = []testData{
	testData{
		s:        "()())()",
		expected: []string{"(())()", "()()()"},
	},
	testData{
		s:        "(a)())()",
		expected: []string{"(a)()()", "(a())()"},
	},
	testData{
		s:        ")(",
		expected: []string{""},
	},
	testData{
		s:        "x(",
		expected: []string{"x"},
	},
	testData{
		s:        ")))((()((s((((()(((",
		expected: []string{"()(s)", "(()s)", "()s()"},
	},
	testData{
		s:        "(()((s(()",
		expected: []string{"()(s)", "(()s)", "()s()"},
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, removeInvalidParentheses(item.s))
	}
}

// func TestIsValid(t *testing.T) {
// 	type data struct {
// 		s        string
// 		expected bool
// 	}

// 	d := []data{
// 		data{
// 			s:        "((()))",
// 			expected: true,
// 		},
// 		data{
// 			s:        "(aaaa(()))",
// 			expected: true,
// 		},
// 		data{
// 			s:        "(aaaa(())",
// 			expected: false,
// 		},
// 		data{
// 			s:        "()())()",
// 			expected: false,
// 		},
// 		data{
// 			s:        "()()()",
// 			expected: true,
// 		},
// 	}
// 	for _, item := range d {
// 		assert.Equal(t, item.expected, isValid(item.s))
// 	}
// }
