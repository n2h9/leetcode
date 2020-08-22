package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	A := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	expected := 6

	assert.Equal(t, expected, maxSubArray(A))
}

func TestSol1(t *testing.T) {
	A := []int{-2, -7, -3, -1, -5}
	expected := -1

	assert.Equal(t, expected, maxSubArray(A))
}

func TestSol2(t *testing.T) {
	A := []int{5, 5, 5, 5, 5}
	expected := 25

	assert.Equal(t, expected, maxSubArray(A))
}

func TestSol3(t *testing.T) {
	A := []int{-1, -1, -1, -1}
	expected := -1

	assert.Equal(t, expected, maxSubArray(A))
}

func TestSol4(t *testing.T) {
	A := []int{-3, 2, -4, 4, -1, -1, -1, 2, 1, 1, -5, 4}
	expected := 5

	assert.Equal(t, expected, maxSubArray(A))
}

func TestSol5(t *testing.T) {
	A := []int{-3, 2, -4, 4, -1, -1, -1, -1, 2, 1, 1, -5, 4}
	expected := 4

	assert.Equal(t, expected, maxSubArray(A))
}
