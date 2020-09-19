package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	n             int
	prerequisites [][]int
	queries       [][]int

	expected []bool
}

var data []testData = []testData{
	// n = 2, prerequisites = [[1,0]], queries = [[0,1],[1,0]]
	testData{
		n:             2,
		prerequisites: [][]int{[]int{1, 0}},
		queries:       [][]int{[]int{0, 1}, []int{1, 0}},

		expected: []bool{false, true},
	},
	// n = 2, prerequisites = [], queries = [[1,0],[0,1]]
	testData{
		n:             2,
		prerequisites: [][]int{},
		queries:       [][]int{[]int{1, 0}, []int{0, 1}},

		expected: []bool{false, false},
	},
	//Input: n = 3, prerequisites = [[1,2],[1,0],[2,0]], queries = [[1,0],[1,2]]
	// Output: [true,true]
	testData{
		n:             3,
		prerequisites: [][]int{[]int{1, 2}, []int{1, 0}, []int{2, 0}},
		queries:       [][]int{[]int{1, 0}, []int{1, 2}},

		expected: []bool{true, true},
	},
	// 	Input: n = 3, prerequisites = [[1,0],[2,0]], queries = [[0,1],[2,0]]
	// Output: [false,true]
	testData{
		n:             3,
		prerequisites: [][]int{[]int{1, 0}, []int{2, 0}},
		queries:       [][]int{[]int{0, 1}, []int{2, 0}},

		expected: []bool{false, true},
	},
	//Input: n = 5, prerequisites = [[0,1],[1,2],[2,3],[3,4]], queries = [[0,4],[4,0],[1,3],[3,0]]
	// Output: [true,false,true,false]
	testData{
		n:             5,
		prerequisites: [][]int{[]int{0, 1}, []int{1, 2}, []int{2, 3}, []int{3, 4}},
		queries:       [][]int{[]int{0, 4}, []int{4, 0}, []int{1, 3}, []int{3, 0}},

		expected: []bool{true, false, true, false},
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, checkIfPrerequisite(item.n, item.prerequisites, item.queries))
	}
}
