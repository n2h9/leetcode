package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	nums []int

	expected [][]int
}

var data []testData = []testData{
	testData{
		nums: []int{-1, 0, 1, 2, -1, -4},

		expected: [][]int{{-1, 0, 1}, {-1, -1, 2}},
	},
	testData{
		nums: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, -2, -2},

		expected: [][]int{{-2, 1, 1}},
	},
	testData{
		nums: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

		expected: [][]int{{0, 0, 0}},
	},
	// testData{
	// 	nums: []int{39347, 50937, -13916, 81279, -73332, -17101, 5173, -4330, -58673, -59218, 28773, -52855, -97173, -4377, 27742, -77469, 13983, -47372, 3090, 9323, -81881, -85777, 23047, 11675, 30588, 95136, 57617, -26141, 44158, -60952, 86401, 84071, 34416, -10312, 39502, 86489, 48534, 61226, -35916, -98449, 20337, -9843, 26350, -16717, 59869, 13562, 79061, -55020, -77138, -9828, -1321, -26685, 73017, 96166, 40760, -15721, -66133, 14899, -49794, -90899, 85449, -42009, -39626, 77740, 64516, 71013, 38569, 42611, -99448, -50076, 95448, 42354, -7800, 57097, -51638, -36490, 71484, -14595, 62758, 81835, 74486, 40954, -54083, -8596, -46338, -88893, -813, -19297, -12415, -79348, -80923, -69082, 89687, 43832, -19516, 65180, -47885, 32672, -18936, -15895, 71605, 39453, 7332, -75059, 76558, -79011, -8423, 80208, 2108, 13760, -31777, -98401, -95325, -95550, -97554, 86822, -84023, -11348, -34956, 24007, -25885, -72368, 62174, 32617, -44278, -63120, 55739, -85206, -87409, 33736, 31030, -7338, 86557, -61225, 4597, 50387, -38964, -35201, -4479, 13987, -1215, 90117, -22902, 92835, 56224, -22149, 27779, -12834, 14465, -91088, 78358, -97254, -20023, 91122, -33351, -65222, -85603, 45851, 57996, -83370, 57330, -25215, -33524, -81872, -13481, -49274, -8883, -93148, 82211, -7071, 33398, -19182, 81142, 10184, 40332, 73091, -99926, 49552, -80383, 18121, -29135, -80616, -82863, -82325, -47337, 65493, -50959, 11223, 95416, 22150, -30435, 18382, 98801, -55706, -6967, 66534, -63047, 24167, -98079, -22344, 94036, 14578, -38169, 22514, -88854, 94746, 6518, -54730, -72227, 50943, 71264, -35215, 22582, 24429, 21273, -98198, -50028, -83447, 14789, 76993, 82225, 30375, 66997, -7914, -71596, -40898, 94510, 96118, -50510, -29842, 53326, 94254, -24817, 71371, 60201, 40399, 37914, -33667, 838, -43147, -73545, 6814, -46102, -70359, -37433, -36856, 34161, -63276, 21202, 88253, 52920, -48628, 92579, 32343, 25169, 13717, -36322, 7743, -46118, 63527, 53476, -24464, 12759, -81691, 63294, -92078, -75257, 38747, 89320, 41490, -78248, -3422, 81903, -28826, 70639, -37873, -9993, 41710, 33850, 82934, -87427, -42838, -19399, 29480, 5490, -18527, 34705, -17397, 66838, -19782, 50270, 29869, -630, 51664, -79123, 74639, -65177, -8598, 68896, -32956, -14322, 82876, -71376, 2509, -32363, -1977, -2029, -95049, 20369, -31164, -62279, 80760, -23984, 33066, -8841, -97855, 99206, 5051, -60861, 99452, 62389, 89801, 11779, -85450, -87203, -65988, -26524, -47001, -77801, -3294, 8130, -7756, -62173, -92829, -71649, 34714, -86263, 85552, -46689, -91549, 68617, -47233, 44597, -38818, 81976, -1511, -28382, -3930, 40404, -77222, -3926, -1301, -36085, -58642, 32967, 52615, -28670, 23939, 41797, 19201, -86306, 8266, 44587, 87119, 73386, 6132, -65189, -52883, 40594, 63224, 72566, 81299, 46831, 98705, -45530, 15458, -95857, 124, -94501, 29945, -79203, 61154, 51845, 81153, 97059, -63084, -71533, -85218, 71537, -56633, -86987, 36313, 32021, 96659, -75817, -20886, -31511, 99578, 58914, -62482, -46584, -96156, 57720, -8488, 32505, -85960, -90629, 34938, 81962, 61969, 46950, 34463, -40556, -80788, -81219, -41738, 75289, -14810, 9421, 49824, 23413, -69793, -81457, -26824, 28655, 1465, 93802, 30528, 78429, 64260, 23212, 3571, 65063, -59660, 5999, 74512, 90902, -17443, 76741, 31190, 76227, 10982, 24253, 78765, 28685, -73127, -19889, -19640, -67521, -20375, -36906, 89495, -50470, -65107, 43787, -71101, 61866, 5161, 58808, 57894, -18857, -49045, 19435, -82338, 24692, -93900, 4913, 12332, 82688, -67185, -71918, -58245, -1275, 54472, 49319, 82643, -21627, -55302, -63867, 45557, 77319, -47480, -76339, 27607, -14231, 66669, -77467, 81927, 64510, 90068, 11664, -30423, -78173, -88474, 39098, 46166, -74692, -97036, 55564, 38651, -98736, 93831, -78123, -91835, 70676, 89240, -80898, 31019, -99313, -80102, -90100, -97900, -5554, 81222, -83076, 70464, 64449, 65026, -41629, 25005, 24979, 65149, 43865, 9316, -32079, 14776, -69695, 97285, -89721, 96222, 14356, 1755, -60328, -46176, -53267, -56112, 51145, 48780, -25580, 97806, 74900, -91579, -85670, 38046, -96052, -7641, -42539, -90928, -32701, -25059, 1559, 9514, -28774, -22688, -20642, 17845, -80474, -82186, 93646, 72633, 41901, -75044, -6333, -49709, -43687, -95663, 1178, -69912, -34708, -65666, -13708, 68600, 48811, 35449, 48745, -20932, 57629, -49601, -18470, -62567, 7293, -82331, -29302, -29989, 50608, -52902, -98105, 42595, 53476, 42842, -19293, 95169, -26528, -25775, 85304, 73706, 7409, -68948, -15528, -61822, -36972, -32664, 39819, 9275, -66341, -41103, 16834, -86010, -58175, 40028, -62859, 21414, 23940, -47267, -24832, -72559, -66312, -57990, -42639, 82600, -5165, -16160, 94992, 70325, 22350, 35705, -6518, 21165, -10810, -33566, 5796, -7654, 64156, -60452, 89687, 67498, -86143, 78579, 48410, 45950, -18119, -99240, 14267, -60759, -34152, -66294, 67815, 10311, 89367, -4829, 39977, 70281, -20286, -37330, -12612, -12807, -46935, 39400, -94728, 69411, -96456, -44278, 37571, 91707, 4665, 53303, 97971, 28007, -91673, 56709, -31022, 67211, -38730, 19393, 89446, -70680, -83254, -80359, -17247, -57153, 31537, 90472, -71594, 74864, 69789, -94381, 9193, 28204, -29971, 65849, 43568, -30581, -94649, 77688, 3974, 9880, 86761, -44286, 30355, 47894, 99578, -70623, 12188, 75106, 40164, -60393, 46040, -72074, -66949, -1006, -3316, -73562, 99673, -75242, 45030, 70974, 50205, -29667, 77880, -70279, 96262, -91429, -47169, -8390, 74662, -21458, 73653, -40194, -59273, -29124, 75171, 64940, -50638, 85709, 88468, -5598, 20066, -77976, 7971, 23349, 57472, -29491, 43984, 94587, -11581, -9020, -62707, 36858, 24663, -36531, 51634, -91681, 79467, -15534, -54654, 83289, -17885, -32421, 68856, -56071, -9349, -22457, -61994, -91130, -10228, -37376, -51534, 50987, 32974, -43827, -39843, -58599, 64523, -35347, 99102, 70781, 53801, -95742, -96026, -64658, 51313, -32069, -78882, -31126, 51839, -38595, -84412, 99828, 60064, -83654, -17261, -49302, -10002, -46797, -15256, -19120, -56172, 56249, 38725, -3380, -96408, 45315, 7055, -32959, -75360, 53169, 65067, 44263, 64459, 85588, 86521, 72040, -89457, -66322, -83542, -74008, 38832, -45431, -15266, 63705, -52536, -11194, -70855, -4803, -43819, -76564, -56198, 14869, -47657, 9059, 60959, 59443, 75432, -89232, -76245, 66948, 8861, 20905, 39466, -6170, 86935, -8876, 97511, 56672, -98546, -58307, -52492, -32564, -47486, -64580, 30223, 67195, 18500, 59187, -34505, 40872, -70459, -49232, -91547, -37739, -83340, -68581, -89862, -52633, 27278, -15970, -32287, -11649, -66387, -69581, 96436, 16984, 48673, -89864, 92123, 37769, 96409, -3215, -23972, 4635, 1014, 82900, -64961, 40292, -78721, -50972, 11630, -97709, -42897, -58162, -87164, 97331, -7332, 81270, 71369, -30586, 19125, -46667, 67708, -42403, 21173, -32027, -92924, -9856, -93578, 78755, 79285, -38743, -88005, -63017, 74680, 48195, -37162, 16762, -3048, -81782, 31927, 15232, 72002, -43756, -78847, 31810, 59372, -54542, -34046, -46228, 30810, -8591, -82994, -63358, 64609, 628, 36976, -13768, -11068, -84453, 51758, -54008, -56289, -69554, -79712, 1563, -46595, 70678, -21542, 76355, -42027, 28826, -32478, 52534, -23904, 31681, -47230, -80791, 86492, -77850, -95302, -45480, -72148, -88013, 5532, -45747, -79725, -80312, 81572, -26097, -60400, 84792, 39336, 4310, 73911, 69510, 13975, 68391, 15308, -81787, 513, -42234, -49907, -43650, 99006, 7452, -89416, 14865, 63334, -73539, 85353, -96455, -13677, 76871, -90671, -81600, 65660, 52296, 69062, 73178, -81365, 21155, -36614, -61142, 22194, 21423, 23175, -59651, 40439, -57476, 60623, 6860, 46770, -2306, 88937, 92681, -31407, -70477, 86478, -92338, -9146, -14942, 70147, -9576, 71947, 60467, 7561, 89876, -87345, -50050, 73793, 13175, -64856, 29791, -74181, -81708, 4745, -94682, 8825, -29914, -26944, 51431, -13430, 11739, -33386, 32950, 70114, -31309, -63330, -58433, -62116, 60729, -15723, -7711, -12903, 37985, -87232, -57999, -57955, 27289, -41340, -55911, 85151, -39781, -67392, 93761, 83075, -56961, 97637, 28277, -215, 91037, 66818, 69982, -26635, 1881, -98027, -77625, 97006, -75290, 44640, 14364, -52076, 72090, -32012, -79434, 48108, -83407, -56055, 29843, -30615, -467, -4010, -22824, 84489, -19864, -53565, -28977, 83263, -27775, 29952, 36774, 74016, 23548, 30320, 51330, 70874, -80941, -94494, -27636, 51152, 19975, 45972, -59544, -8123, -16287, 48942, 78918, 31958, -55706, 60585, 39468, -29098, 37205, -7078, 18616, -64749, 9768, 49252, 69774, 22900, -34544, 54401, -23869, -59951, 54224, 74439, 19542, 9693, -56834, 546, 54976, -411, 2907, 98733, 60600, 28407, -67888, -67032, -81396, 40452, 50515, 34181, 28863, 91550, 38837, -95751, 78366, 24525, 61590, -70477, 54038, 61450, -68725, -32027, 73369, 23013, 62699, 3876, 28354, -53757, -38473, -52709, 40521, 11318, -19856, 18105, 29796, -24216, 8050, 47854, 39703, -68569, 84565, -22389, -39459, 45746, 28869, 58601, -26707, -89254, -52050, 15712, 96367, -11535, -62761, 14431, -75124, -31454, -59485, -57004, 19452, 39564, 79988, -56628, -29037, -68245, 63724, -83075, -72257, 43168, 3976, 69008, 52737, -69731, -61810, -71344, 17898, -35098, 42849, 50295, -14065, -2841, 44285, -45803, -71371, -26973, -19158, 48304, -59264, 40633, 78167, -42385, -52276, 22033, 22316, -56140, 78704, -5538, 66169, 99758, -62242, 11356, -44325, -20294, 75495, 10664, -57244, -25365, 29584, -12557, -29912, 92774, 54424, -71512, -56215, -24423, 45533, 52966, 21553, -95110, -75889, -26454, 48951, -66792, 95279, 74741, 41411, -11001, 6658, 34677, -11521, 34662, 45043, 67410, -48963, 13343, -3357, 23006, 94416, -95666, -16610, -7460, -69473, -20368, -65247, 2883, 81794, 92797, -33077, -42545, -35410, -84072, 32486, 69168, -77958, 17706, -22966, 54293, -19104, 50955, 71373, -81340, 35399, 67508, -31372, -94775, -69101, -95084, 36749, -44157, -13909, 39873, -52016, 34777, -92290, -14329, -91303, 55112, 14297, -20750, -14629, -51438, 30404, -69346, 65384, -91076, -71053, -69998, -19627, -78897, -16053, -20443, -89459, 32280, 67565, 58293, -4856, 68540, 82253, 26302, 93657, -85916, 48393, -48966, 91378, -22077, 22392, 4849, -78629, -51696, 52991, -91352, -68970, 90940, 67595, -62546, -70014, 32232, -31278, -63778, -99460, 68248, -64061, -38311, 21792, 90490, 12405, 58115, 75976, -45561, -1822, 52608, -50374, 10348, -97348, 39521, -75416, 86129, -67898, -53328, -75378, -51776, -39339, -47446, 89484, 51704, 29021, 3045, 5387, 70536, -10889, -40482, -30013, 53067, -57983, 30812, -37357, -27303, 81170, 31342, 29371, -46877, 99631, 77302, 77171, -14503, -65322, -66934, 24115, -53711, -27823, -79950, 45547, -65576, -94098, 60525, 50585, 62464, 65908, 55444, -74308, -85786, 96211, -64638, -52798, 56961, 45562, 51840, 78028, 75888, -17742, -99849, -22274, 94597, 26088, -39096, 39963, 10184, -50858, 88528, -13165, -17153, -503, 71016, -77121, 65426, -4468, -56771, -67837, -86595, 38986, 14411, -39226, -62603, 41737, 69195, 76940, -88866, -60867, -27591, -69297, 59018, -51596, -90722, -7998, -58547, -92169, -24209, -35154, 80498, -24364, 76893, 97455, 51979, 41516, -92882, 3610, -37758, 25766, -15410, 93950, -66528, -7289, 64848, -50453, 70450, -54011, -28124, 8383, 58721, 71481, -22992, -68944, 29307, -68938, -12625, 6924, 67140, 26674, -14071, -52249, 55793, 70927, -85502, 64598, 54601, -74354, -74320, -29071, 69490, -89570, 9026, 39572, 60793, -40805, 7783, 59128, 20101, 25893, -63142, 77888, 53047, 99912, 10350, 33380, 73086, -72789, -30854, -63916, 96329, 78884, 85214, 76395, 70345, -85189, -10247, -45880, 92972, 74364, 39494, 32969, 40298, -93664, 82536, -5687, 10435, 60127, -6938, 90400, 8172, -20255, 76655, -88804, 44976, -28543, -2485, -41544, -79697, 75738, -31617, -94971, -1517, 94294, 43780, -94116, 77797, -58243, 8081, 34482, 11563, 3591, 44971, -56079, -35182, -46702, -56901, -91223, -72035, 41647, -59291, -27127, 32437, 79182, -20002, 67962, 23713, -64140, 55631, 27153, -92592, -96225, 20815, -32497, -49404, 86923, -50882, -2676, 27103, 73626, -84085, 95417, -87473, 78481, 63560, -35211, -8633, 82387, -666, -49412, -34099, 28464, 65055, 46906, 41605, -37156, -30822, -4719, -45698, -73947, 62885, -11176, -49744, -45946, 16530, -33569, -29155, -91154, -8227, -94535, 70697, -22030, -53864, -50160, -43764, -20880, 7438, -48543, -73018, -85639, -11802, 19850, -83270, -80098, -31117, 7645, 37098, -52379, 6934, -81406, 28682, -67357, -77869, -69937, -17416, -78925, -96038, -32115, 70161, -65110, 46893, 49972, 99695, -44156, -11209, -80785, 30088, -64775, 60085, 79135, -54673, -32876, 75473, -66446, -69902, -50550, -59989, -93011, -26358, -85501, 57168, 32213, 54363, -84075, 51027, -24206, 55312, 7584, 25922, 67625, 2473, 37636, 53936, 21758, 92684, 49936, -49848, 99059, 55218, 28961, 50956, 43749, -63536, -20672, 14550, 58392, -55779, -95741, 37357, 67229, -55096, 71823, 7440, 50032, -93842, -16456, -60324, -97985, 84064, -63826, -81328, 65320, 3753, -55049, -8933, -77692, -97350, 53631, 15466, 57948, -11582, -99127, 72683, 33215, 84749, -83025, 91585, 71007, 91555, 75852, -1962, 26753, -25935, -8282, 83909, -43735, -60298, -32848, 48157, 81757, -6814, 49712, -27006, 6364, -24105, -80049, 23925, 84707, 25944, 97834, -44586, -44212, 68523, -66304, 96585, 29550, -61159, 85755, -85913, 34075, 627, -81761, 49258, 5846, 95703, -37064, -28274, 92969, 42326, -10851, -19150, -24123, -32317, -2183, -25569, -29315, -94438, -49381, -18432, 96350, 84916, 65823, -39089, 5183, 83920, 98867, -69104, 52969, -72734, 39763, 76463, -87311, 20569, -26164, 92195, 25945, 31182, -92886, -21572, 51072, -80771, -88199, 21078, 47312, -44428, 25660, 20275, -92258, -20939, 98917, 94689, -76241, -51644, 47622, -50611, 92352, -57836, -44836, 84353, 6588, 58345, 89949, 59109, -88513, -33835, -54270, 81008, -86906, -45563, 6458, -88952, -40695, 52284, 17603, 14364, 65317, 6781, 39123, -24243, -99524, -58832, -64288, 30289, 25729, -83089, 57526, -40806, 1417, 81944, 99354, -38838, -56380, 15099, 3539, -10573, 26648, 70370, -9604, 66880, -75448, -38979, -77355, 49381, 93966, 22494, -2921, 71238, -21943, -49149, 23319, -47235, 69529, 59011, -22337, 26045, 517, -71597, 19977, 57661, 43337, 10370, -46967, -27978, -1297, -8699, 69441, -11051, -20774, 80267, -64930, 41132, -86726, -54996, -19428, 74572, 86051, 53694, 48608, -7373, -70041, 32623, -41117, 12047, 39213, 46141, 89177, -51915, -41748, -87171, -52961, 31775, -750, -12903, 85341, 58310, -88875, -61503, 32097, 99876, -7662, -78915, -80506, -38530, -77739, -86309, 7853, -66163, 60858, 94356, -67571, 24989, 24968, -42946, -62777, -87430, 38731, 37747, 15851, 81868, -78361, -85497, 65068, -811, -55891, -84439, -36263, 41318, 97342, -20570, -67865, 59778, -75212, -78421, -7703, -10231, 34639, 15013, 4560, 88604, 36922, 68346, 74193, 89360, -26647, 63595, -34951, -61060, -86158, 85612, -5374, 99924, 15527, 28994, -77651, -49143, -50486, 87661, -78721, 76399, 66658, -28341, -69971, -48973, -92743, 2218, 79020, -89449, -72644, 52832, 52350, 61669, 64225, -32232, -61553, -83936, 88031, -75916, 77350, 43010, 10602, -92110, 64580, 7857, 699, 76958, -82288, -76831, -12801, -6958, -2450, 28943, -45426, -92037, -77173, -70373, 46470, -5452, -4956, -5451, -3716, -61993, -36654, 48128, 94566, -4361, -65390, 40289, 54338, -94426, 42677, 48787, -26460, 86515, 90067, 92652, -26642, 75381, 38993, -50262, -29370, -19774, -28024, -64524, 84503, 97440, -90902, -53793, 72749, -90183, 37565, -17572, -28039, -36459, -52226, -13627, -21858, 52605, -27833, -72029, 31934, 58647, 45130, 79585, -10669, 2011, -31877, 94118, -81819, 24736, 66150, 51119, 73084, 22778, -24169, -51339, 93050, 90554, -58634, 42762, -81273, 87275, 90558, -74212, 96804, -91180, 35366, 36206, 70677, -34066, -96524, 67947, 32885, 79198, 3648, 21839, 15455, 5148, -50359, -63656, -47321, 67957, -783, -2805, 86896, 63363, 58767, 51258, 24421, 21146, -33428, -92739, -4728, 73852, 9272, 44211, -68125, -17580, 55886, -25077, -42911, -54304, 98524, 31526, -79566, -47025, 46603, -31004, 69597, 34514, 49620, 5501, 84665, -11582, -79092, 33810, -1527, 39496, -24925, 52390, -68094, 4372, 72722, -20942, 25954, 42838, 48727, -46581, 16147, -33900, -36805, 90297, -89705, -98406, -58213, 16215, -92925, -28048, -6963, -21076, 42667, 10689, 35972, 47642, 40097, 71610, -10929, 27568, -31184, -593, 33695, 97961, 21923, 11262, -90472, 2907, -91981, 61016, -76152, -21648, -73749, -14079, -70242, 5321, 2977, -78549, -9654, 19195, 23492, 37678, -52507, 17744, 71869, -72441, -61266, -35072, 3611, 49817, -98415, 64614, 35648, 56277, -67150, 58435, 10026, -29372, -6666, 45762, -80023, -1420, 49266, -86109, 47173, 11070, -32924, 25597, -16868, -19968, -60772, -40177, -89799, 58824, -50186, -78421, 25342, -42429, -73686, -48078, -31246, -9569, -49011, -73134, -27249, -69927, 45630, -71551, -34543, -70910, 11513, -42107, -97821, 95793, 84140, -10763, -95983, 19929, 14012, 46659, 91357, 57600, 59892, -60904, -45596, 97028, 98992, -29778, 92047, 76883, 45459, -98962, -50502, 4037, 51502, -33366, 14977, 11149, 26437, -89595, 9311, -74104, 87779, -7684, 89399, 24982, -38461, 86072, 7470, -71235, 84921, -27714, -68044, 92599, 36663, -92487, 73502, 24389, 12019, -91923, -82952, 44944, 70682, -65298, -77658, -74049, -9816, -24437, -1184, -84748, -39219, 32617, -18288, 60415, -44758, -30706, 24560, 81363, -31365, 34047, -13945, -28898, -32215, -93767, -47158, 62945, 82410, -53778, 98575, 63963, -71802, -8492, 22263, 83594, -72671, 17125, -52024, 49349, -98875, -60870, 17186, 63997, -45870, 7818, -70161, 52446, -41196, -33391, 14602, 79997, -45218, 52152, 72270, -15873, -60021, 21658, 95459, -60581, -55651, -53032, -17884, -92140, 26825, 73986, -66657, 46654, 31839, -53107, 39297, 46942, -92091, -43664, 69132, 18513, -61213, 82503, 12171, -99120, -81708, -4928, 25142, 38145, 24627, 15993, -85551, -79475, -86348, -22783, 85194, 55122, -12258, -61313, 12737, 74319, 54715, 6372, -15032, -52438, -82383, -46885, 39643, 23517, -8688, 97424, 38905, 11271, -74359, -37942, -90691, -88805, 90187, 93535, -31450, 68476, -18713, 32484, 24347, 26781, 22155, 72660, 45713, 71708, 97735, -7006, 70691, -16473, 78504, 78097, -98085, -83732, 44320, -82919, -23278, 47447, 43685, 24248, 95821, -91228, -32670, -53078, 41606, 45171, 64140, 23522, 26776, 31453, -52860, 13506, -52179, -81187, 49161, 80937, 97612, -30852, 82264, -52875, 94131, 32422, 35217, 95225, -31035, -94969, 5481, -7798, 30560, -89112, 8043, -41717, 32119, 37186, 53440, -68978, -90709, 10038, 37168, -13831, 84179, 97144, -94885, -68630, -43856, 73431, -81323, -9290, -73782, -13173, 564, -54878, -72041, -92591, 27239, 12678, -92361, -96458, 17661, -21207, -52462, -51267, -85787, 62996, 50541, 43054, 74258, -16838, 61758, -71861, -63397, 53655, 42558, -94291, 46873, 88928, -75584, -55197, -34718, -52627, -43364, 67309, -67808, 3788, -43353, -28168, 14356, 35721, -90641, -97741, -21636, -75674, -85422, -8874, -72974, -21824, 98779, -65367, -25414, 2844, -58395, 70544, -35908, -28527, 83677, 88645, -89860, 30404, 83453, 66095, 96655, 63936, -17952, -30835, -57705, 32757, 51451, 82083, 52912, 86226, -86491, -63284, 93157, 42630, 88573, 26451, -2779, -13869, -80365, 79389, 28838, 3939, -35951, 69897, -20512, -61395, -89988, -56731, 96372, -13016, 15166, 73215, -79145, -28990, -32841, -18723, -24852, -45442, -49939, -30351, 19471, 5895, -33346, -52544, 92207, 8745, -79781, -36863, 72225, 93169, -65667, -27889, 20472, 60473, 27632, -44904, -3321, 30621, -16217, -95374, 94212, -9748, -5189, -44248, 76763, 22206, 77330, 60861, -79816, -9790, 82341, -40995, -18858, -70926, -20254, 16576, 93891, 39546, 97404, -87076, 75739, -35545, -74622, 64426, -47656, 51448, -20710, 45542, 99389, -82630, 7162, 40903, 20846, 29363, 30096, 86535, -67839, 31256, 26515, -7234, -78456, -90765, 15231, -70531, -15356, 11977, -35362, 82483, -12396, -89438, 29939, 31484, -60446, -35609, -91743, 79203, -55313, 5353, 28180, 47624, 36233, -71073, 99022, -58426, -37546, -1427, 78873, 90172, -11964, -23961, -75402, -59224, -34470, -86917, 79518, -95179, -96676, -66506, -57555, 16072, -33960, 33700, -13218, -78888, -4778, 75222, -64488, 23208, -8263, -59565, -62126, 89248, 29498, -26623, 33157, -84761, 14147, -75736, -74616, -97340, -67456, 51416, -41201, 63210, 62450, 40455, 9739, 71271, -78882, 18538, -24324, -38099, 88090, 82762, 97581, 96986, -26781, 83892, -86744, 50340, -79495, -96894, -68860, -62931, -38101, -61449, -60181, -86775, -43536, -63696, 9560, 95903, 28321, -94566, -75290, 71217, 39500, -32789, -30664, -30774, 92168, 45473, -44100, -45923, 95610, -87953, -78084, 66349, 59186, 98349, -13449, -20165, -86021, 94628, -88642, -10942, 8295, -92311, -39703, -64976, 90568, 22819, 13118, 26947, 23763, 4704, -58563, -23056, -245, 25165, -65178, -82375, -55966, -7560, 98878, -91468, 88068, -50681, -85126, -78231, -27844, 79838, -39550, 94141, -35848, 80056, -85133, -58788, -7792, -64337, -61352, -47484, 47915, -48096, -20006, -73357, -9438, -77351, -41835, -60031, -77840, -26939, -88485, 74537, 74357, 86053, -79631, 2033, -23671, 72344, -92495, 57145, 79278, -76884, 51834, 22983, 10239, -28910, 76497, -79957, -59153, -68860, 54171, -27938, -79397, 54479, -14142, -321, -65619, -92012, 4, 30147, 27131, -96688, 46001, -19379, -97163, -2826, 54575, -20003, 4118, -53351, -75159, -35093, -2233, -23501, 51655, -53498, 77028, -59886, 8571, 78154, -79859, 23864, 51805, 65183, 7771, -91479, 9815, -36634, -2644, 17022, 54106, -34245, -63457, 4727, 13808, -80215, -82517, 92974, -23705, -72111, 82417, -77286, 4604, -12430, 11940, 47004, 63759, 10257, 20467, -3902, 80533, -38176, -99550, -45065, 36771, -94044, 14738, -53204, -48074, 90657, 13904, 78633, 93554, -14078, 31003, 50528, 26221, 93905, 6230, 19107, -53248, 30526, 17082, 82926, -42946, -96897, -87392, 52534, 82177, 16155, 13186, 78816, 18754, 86620, 79907, -60960, -66372, 10796, 76217, 66016, 52608, 21983, 13446, 97234, 80186, -65407, 43325, -14874, -29573, -21728, 64621, -1911, 40782, -75021, -83204, -12594, 45843, 38951, 27012, 34136, -81313, -41070, 22997, 26851, 29391, 32467, -84319, 59438, 15492, 46284, 98870, -78625, 75820, -84605, 84265, -70973, -74774, -84370, 24470, -39460, 50282, -20013, 88548, 20319, 94898, 69961, 19526, -53076, -70510, -39769, 49089, 63380, 40215, 84507, 16108, 27315, 68358, -64329, 34435, -26561, 62082, 93132, 19747, 38236, -50858, -79606, -22589, 61198, -25324, -1098, -57509, 89496, 1421, -55025, -28506, -39621, -46550, 25038, 13379},

	// 	expected: [][]int{},
	// },
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, threeSum(item.nums))
	}
}