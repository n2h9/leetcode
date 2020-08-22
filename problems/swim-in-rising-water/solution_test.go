package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	grid := [][]int{
		[]int{0, 2},
		[]int{1, 3},
	}
	expected := 3

	assert.Equal(t, expected, swimInWater(grid))
}

func TestSol1(t *testing.T) {
	grid := [][]int{
		[]int{0, 1, 2, 3, 4},
		[]int{24, 23, 22, 21, 5},
		[]int{12, 13, 14, 15, 16},
		[]int{11, 17, 18, 19, 20},
		[]int{10, 9, 8, 7, 6},
	}
	expected := 16

	assert.Equal(t, expected, swimInWater(grid))
}

func TestSol2(t *testing.T) {
	grid := [][]int{
		[]int{20, 1, 2, 3, 4},
		[]int{24, 23, 22, 21, 5},
		[]int{12, 13, 14, 15, 16},
		[]int{11, 17, 18, 19, 0},
		[]int{10, 9, 8, 7, 6},
	}
	expected := 20

	assert.Equal(t, expected, swimInWater(grid))
}

func TestSol3(t *testing.T) {
	grid := [][]int{
		[]int{0, 6, 3, 2, 1},
		[]int{4, 7, 8, 9, 10},
		[]int{5, 23, 14, 15, 11},
		[]int{24, 17, 18, 16, 12},
		[]int{10, 9, 8, 7, 13},
	}
	expected := 13

	assert.Equal(t, expected, swimInWater(grid))
}
