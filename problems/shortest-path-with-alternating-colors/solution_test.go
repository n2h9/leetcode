package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	n := 3
	red_edges := [][]int{
		[]int{0, 1},
		[]int{1, 2},
	}
	blue_edges := [][]int{}
	expected := []int{0, 1, -1}

	assert.Equal(t, expected, shortestAlternatingPaths(n, red_edges, blue_edges))
}

func TestSol1(t *testing.T) {
	n := 3
	red_edges := [][]int{
		[]int{0, 1},
	}
	blue_edges := [][]int{
		[]int{2, 1},
	}
	expected := []int{0, 1, -1}

	assert.Equal(t, expected, shortestAlternatingPaths(n, red_edges, blue_edges))
}

func TestSol2(t *testing.T) {
	n := 3
	red_edges := [][]int{
		[]int{1, 0},
	}
	blue_edges := [][]int{
		[]int{2, 1},
	}
	expected := []int{0, -1, -1}

	assert.Equal(t, expected, shortestAlternatingPaths(n, red_edges, blue_edges))
}

func TestSol3(t *testing.T) {
	n := 3
	red_edges := [][]int{
		[]int{0, 1},
	}
	blue_edges := [][]int{
		[]int{1, 2},
	}
	expected := []int{0, 1, 2}

	assert.Equal(t, expected, shortestAlternatingPaths(n, red_edges, blue_edges))
}

func TestSol4(t *testing.T) {
	n := 3
	red_edges := [][]int{
		[]int{0, 1},
		[]int{0, 2},
	}
	blue_edges := [][]int{
		[]int{1, 0},
	}
	expected := []int{0, 1, 1}

	assert.Equal(t, expected, shortestAlternatingPaths(n, red_edges, blue_edges))
}

func TestSol5(t *testing.T) {
	n := 5
	red_edges := [][]int{
		[]int{2, 2},
		[]int{0, 1},
		[]int{0, 3},
		[]int{0, 0},
		[]int{0, 4},
		[]int{2, 1},
		[]int{2, 0},
		[]int{1, 4},
		[]int{3, 4},
	}
	blue_edges := [][]int{
		[]int{1, 3},
		[]int{0, 0},
		[]int{0, 3},
		[]int{4, 2},
		[]int{1, 0},
	}
	expected := []int{0, 1, 2, 1, 1}

	assert.Equal(t, expected, shortestAlternatingPaths(n, red_edges, blue_edges))
}
