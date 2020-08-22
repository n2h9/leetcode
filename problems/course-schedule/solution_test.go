package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	numCourses := 2
	prerequisites := [][]int{[]int{1, 0}}
	expected := true

	assert.Equal(t, expected, canFinish(numCourses, prerequisites))
}

func TestSol1(t *testing.T) {
	numCourses := 2
	prerequisites := [][]int{[]int{1, 0}, []int{0, 1}}
	expected := false

	assert.Equal(t, expected, canFinish(numCourses, prerequisites))
}
