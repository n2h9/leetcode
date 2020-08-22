package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	numCourses := 2
	prerequisites := [][]int{
		[]int{1, 0},
	}

	expected := []int{0, 1}

	assert.Equal(t, expected, findOrder(numCourses, prerequisites))
}

func TestSol1(t *testing.T) {
	numCourses := 4
	prerequisites := [][]int{
		[]int{1, 0},
		[]int{2, 0},
		[]int{3, 1},
		[]int{3, 2},
	}

	expected := []int{0, 1, 2, 3}

	assert.Equal(t, expected, findOrder(numCourses, prerequisites))
}
