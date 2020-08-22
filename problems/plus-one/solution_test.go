package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	A := []int{1, 2, 3}
	expected := []int{1, 2, 4}

	assert.Equal(t, expected, plusOne(A))
}

func TestSol1(t *testing.T) {
	A := []int{4, 3, 2, 1}
	expected := []int{4, 3, 2, 2}

	assert.Equal(t, expected, plusOne(A))
}

func TestSol2(t *testing.T) {
	A := []int{0}
	expected := []int{1}

	assert.Equal(t, expected, plusOne(A))
}

func TestSol3(t *testing.T) {
	A := []int{1, 9}
	expected := []int{2, 0}

	assert.Equal(t, expected, plusOne(A))
}

func TestSol4(t *testing.T) {
	A := []int{9}
	expected := []int{1, 0}

	assert.Equal(t, expected, plusOne(A))
}

func TestSol5(t *testing.T) {
	A := []int{9, 9, 9, 9, 9}
	expected := []int{1, 0, 0, 0, 0, 0}

	assert.Equal(t, expected, plusOne(A))
}
