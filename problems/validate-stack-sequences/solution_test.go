package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	A := []int{1, 2, 3, 4, 5}
	B := []int{4, 5, 3, 2, 1}
	expected := true
	assert.Equal(t, expected, validateStackSequences(A, B))
}

func TestSol1(t *testing.T) {
	A := []int{1, 2, 3, 4, 5}
	B := []int{4, 3, 5, 1, 2}
	expected := false
	assert.Equal(t, expected, validateStackSequences(A, B))
}

func TestSol2(t *testing.T) {
	A := []int{}
	B := []int{}
	expected := true
	assert.Equal(t, expected, validateStackSequences(A, B))
}

func TestSol3(t *testing.T) {
	A := []int{1, 2, 3, 4, 5, 6}
	B := []int{6, 5, 4, 3, 2, 1}
	expected := true
	assert.Equal(t, expected, validateStackSequences(A, B))
}

func TestSol4(t *testing.T) {
	A := []int{1, 2, 3, 4, 5, 6}
	B := []int{6, 4, 5, 3, 2, 1}
	expected := false
	assert.Equal(t, expected, validateStackSequences(A, B))
}

func TestSol5(t *testing.T) {
	A := []int{1, 2, 3, 4, 5, 6}
	B := []int{1, 2, 3, 4, 5, 6}
	expected := true
	assert.Equal(t, expected, validateStackSequences(A, B))
}

func TestSol6(t *testing.T) {
	A := []int{1, 2, 3, 4, 5, 6}
	B := []int{3, 2, 1, 4, 6, 5}
	expected := true
	assert.Equal(t, expected, validateStackSequences(A, B))
}
