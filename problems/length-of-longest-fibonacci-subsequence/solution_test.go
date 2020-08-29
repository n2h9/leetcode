package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	A := []int{1, 2, 3, 4, 5, 6, 7, 8}
	expected := 5

	assert.Equal(t, expected, lenLongestFibSubseq(A))
}

func TestSol1(t *testing.T) {
	A := []int{1, 3, 7, 11, 12, 14, 18}
	expected := 3

	assert.Equal(t, expected, lenLongestFibSubseq(A))
}

func TestSol2(t *testing.T) {
	A := []int{3, 5, 7, 9, 12}
	expected := 3

	assert.Equal(t, expected, lenLongestFibSubseq(A))
}

func TestSol3(t *testing.T) {
	A := []int{3, 5, 7, 9, 12, 19}
	expected := 4

	assert.Equal(t, expected, lenLongestFibSubseq(A))
}

func TestSol4(t *testing.T) {
	A := []int{3, 5, 7, 9, 12, 19, 21, 33}
	expected := 5

	assert.Equal(t, expected, lenLongestFibSubseq(A))
}

func TestSol5(t *testing.T) {
	A := []int{1, 11, 13, 100, 1000, 10000, 11001}
	expected := 0

	assert.Equal(t, expected, lenLongestFibSubseq(A))
}

func TestSol6(t *testing.T) {
	A := []int{1, 2, 3}
	expected := 3

	assert.Equal(t, expected, lenLongestFibSubseq(A))
}

func TestSol7(t *testing.T) {
	A := []int{1, 2, 4}
	expected := 0

	assert.Equal(t, expected, lenLongestFibSubseq(A))
}
