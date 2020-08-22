package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	edges := [][]int{
		[]int{1, 2},
		[]int{1, 3},
		[]int{2, 3},
	}
	expected := []int{2, 3}

	assert.Equal(t, expected, findRedundantDirectedConnection(edges))
}

func TestSol1(t *testing.T) {
	edges := [][]int{
		[]int{1, 2},
		[]int{2, 3},
		[]int{3, 4},
		[]int{4, 1},
		[]int{1, 5},
	}
	expected := []int{4, 1}

	assert.Equal(t, expected, findRedundantDirectedConnection(edges))
}
