package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	n     int
	edges [][]int

	expected [][]int
}

var data []testData = []testData{
	testData{
		n:     8,
		edges: [][]int{{0, 3}, {0, 4}, {1, 3}, {2, 4}, {2, 7}, {3, 5}, {3, 6}, {3, 7}, {4, 6}},

		expected: [][]int{{}, {}, {}, {0, 1}, {0, 2}, {0, 1, 3}, {0, 1, 2, 3, 4}, {0, 1, 2, 3}},
	},

	testData{
		n:     5,
		edges: [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},

		expected: [][]int{{}, {0}, {0, 1}, {0, 1, 2}, {0, 1, 2, 3}},
	},

	testData{
		n:     5,
		edges: [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},

		expected: [][]int{{}, {0}, {0, 1}, {0, 1, 2}, {0, 1, 2, 3}},
	},

	testData{
		n:     4,
		edges: [][]int{},

		expected: [][]int{{}, {}, {}, {}},
	},

	testData{
		n:     5,
		edges: [][]int{{0, 3}, {1, 2}, {2, 3}, {3, 4}},

		expected: [][]int{{}, {}, {1}, {0, 1, 2}, {0, 1, 2, 3}},
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, getAncestors(item.n, item.edges))
	}
}
