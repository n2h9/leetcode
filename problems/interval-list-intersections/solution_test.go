package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	A := [][]int{
		[]int{0, 2},
		[]int{5, 10},
		[]int{13, 23},
		[]int{24, 25},
	}
	B := [][]int{
		[]int{1, 5},
		[]int{8, 12},
		[]int{15, 24},
		[]int{25, 26},
	}
	expected := [][]int{
		[]int{1, 2},
		[]int{5, 5},
		[]int{8, 10},
		[]int{15, 23},
		[]int{24, 24},
		[]int{25, 25},
	}
	assert.Equal(t, expected, intervalIntersection(A, B))
}

func TestSol1(t *testing.T) {
	A := [][]int{
		[]int{0, 2},
		[]int{24, 25},
	}
	B := [][]int{
		[]int{8, 12},
		[]int{15, 23},
	}
	expected := [][]int{}
	assert.Equal(t, expected, intervalIntersection(A, B))
}

func TestSol2(t *testing.T) {
	A := [][]int{
		[]int{24, 25},
	}
	B := [][]int{
		[]int{25, 32},
		[]int{33, 33},
	}
	expected := [][]int{
		[]int{25, 25},
	}
	assert.Equal(t, expected, intervalIntersection(A, B))
}

func TestSol3(t *testing.T) {
	A := [][]int{
		[]int{10, 25},
	}
	B := [][]int{
		[]int{10, 25},
	}
	expected := [][]int{
		[]int{10, 25},
	}
	assert.Equal(t, expected, intervalIntersection(A, B))
}
