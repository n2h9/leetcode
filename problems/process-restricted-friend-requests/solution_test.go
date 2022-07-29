package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	n            int
	restrictions [][]int
	requests     [][]int

	expected []bool
}

var data []testData = []testData{
	testData{
		n:            3,
		restrictions: [][]int{{0, 1}},
		requests:     [][]int{{0, 2}, {2, 1}},
		expected:     []bool{true, false},
	},
	testData{
		n:            3,
		restrictions: [][]int{{0, 1}},
		requests:     [][]int{{1, 2}, {0, 2}},
		expected:     []bool{true, false},
	},
	testData{
		n:            5,
		restrictions: [][]int{{0, 1}, {1, 2}, {2, 3}},
		requests:     [][]int{{0, 4}, {1, 2}, {3, 1}, {3, 4}},
		expected:     []bool{true, false, true, false},
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, friendRequests(item.n, item.restrictions, item.requests))
	}
}

func TestUnionFindForrest(t *testing.T) {
	f := newForest(8)

	assert.Equal(t, 4, f.find(4))

	f.union(3, 4)
	assert.Equal(t, f.find(3), f.find(4))
	assert.NotEqual(t, f.find(2), f.find(4))

	f.union(1, 2)
	assert.Equal(t, f.find(1), f.find(2))
	assert.NotEqual(t, f.find(2), f.find(4))

	f.union(1, 3)
	assert.Equal(t, f.find(1), f.find(3))
	assert.Equal(t, f.find(2), f.find(4))
	assert.NotEqual(t, f.find(4), f.find(7))

}
