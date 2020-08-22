package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	A := [][]int{
		[]int{4, 3, 8, 4},
		[]int{9, 5, 1, 9},
		[]int{2, 7, 6, 2},
	}

	expected := 1
	assert.Equal(t, expected, numMagicSquaresInside(A))
}

func TestSol1(t *testing.T) {
	A := [][]int{
		[]int{5, 5, 5},
		[]int{5, 5, 5},
		[]int{5, 5, 5},
	}

	expected := 0
	assert.Equal(t, expected, numMagicSquaresInside(A))
}
