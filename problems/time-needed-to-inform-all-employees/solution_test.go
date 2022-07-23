package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	n          int
	headID     int
	manager    []int
	informTime []int

	expected int
}

var data []testData = []testData{
	testData{
		n:          1,
		headID:     0,
		manager:    []int{-1},
		informTime: []int{0},
		expected:   0,
	},
	testData{
		n:          6,
		headID:     2,
		manager:    []int{2, 2, -1, 2, 2, 2},
		informTime: []int{0, 0, 1, 0, 0, 0},
		expected:   1,
	},
	testData{
		n:          15,
		headID:     0,
		manager:    []int{-1, 0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6},
		informTime: []int{1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		expected:   3,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, numOfMinutes(item.n, item.headID, item.manager, item.informTime))
	}
}
