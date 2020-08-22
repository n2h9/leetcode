package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	A := [][]int{
		[]int{0, 1},
		[]int{1, 0},
	}
	expected := 1
	assert.Equal(t, expected, shortestBridge(A))
}

func TestSol1(t *testing.T) {
	A := [][]int{
		[]int{0, 1, 0},
		[]int{0, 0, 0},
		[]int{0, 0, 1},
	}
	expected := 2
	assert.Equal(t, expected, shortestBridge(A))
}

func TestSol2(t *testing.T) {
	A := [][]int{
		[]int{1, 1, 1, 1, 1},
		[]int{1, 0, 0, 0, 1},
		[]int{1, 0, 1, 0, 1},
		[]int{1, 0, 0, 0, 1},
		[]int{1, 1, 1, 1, 1},
	}
	expected := 1
	assert.Equal(t, expected, shortestBridge(A))
}

func TestSol3(t *testing.T) {
	A := [][]int{
		[]int{1, 1, 0, 0, 0},
		[]int{1, 1, 0, 0, 0},
		[]int{0, 0, 0, 0, 0},
		[]int{0, 0, 0, 1, 1},
		[]int{0, 0, 0, 1, 1},
	}
	expected := 3
	assert.Equal(t, expected, shortestBridge(A))
}
