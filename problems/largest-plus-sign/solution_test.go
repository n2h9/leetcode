package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	N     int
	mines [][]int

	expected int
}

var data []testData = []testData{
	testData{
		N:        2,
		mines:    [][]int{},
		expected: 1,
	},
	testData{
		N:        5,
		mines:    [][]int{[]int{4, 2}},
		expected: 2,
	},
	testData{
		N:        1,
		mines:    [][]int{[]int{0, 0}},
		expected: 0,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, orderOfLargestPlusSign(item.N, item.mines))
	}
}
