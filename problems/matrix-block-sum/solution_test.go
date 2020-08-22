package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	mat := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	K := 1
	expected := [][]int{
		[]int{12, 21, 16},
		[]int{27, 45, 33},
		[]int{24, 39, 28},
	}

	assert.Equal(t, expected, matrixBlockSum(mat, K))
}

func TestSol1(t *testing.T) {
	mat := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	K := 2
	expected := [][]int{
		[]int{45, 45, 45},
		[]int{45, 45, 45},
		[]int{45, 45, 45},
	}

	assert.Equal(t, expected, matrixBlockSum(mat, K))
}

func TestSol2(t *testing.T) {
	mat := [][]int{
		[]int{1},
	}
	K := 100
	expected := [][]int{
		[]int{1},
	}

	assert.Equal(t, expected, matrixBlockSum(mat, K))
}

func TestSol3(t *testing.T) {
	mat := [][]int{
		[]int{1, 2, 3},
	}
	K := 100
	expected := [][]int{
		[]int{6, 6, 6},
	}

	assert.Equal(t, expected, matrixBlockSum(mat, K))
}

func TestSol4(t *testing.T) {
	mat := [][]int{
		[]int{1},
		[]int{1},
		[]int{1},
	}
	K := 100
	expected := [][]int{
		[]int{3},
		[]int{3},
		[]int{3},
	}

	assert.Equal(t, expected, matrixBlockSum(mat, K))
}
