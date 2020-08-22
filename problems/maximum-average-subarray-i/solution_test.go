package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	A := []int{1, 12, -5, -6, 50, 3}
	B := 4

	expected := 12.75
	assert.Equal(t, expected, findMaxAverage(A, B))
}

func TestSol1(t *testing.T) {
	A := []int{-100, 15, 1000, -600, 150, 334}
	B := 4

	expected := 221.0
	assert.Equal(t, expected, findMaxAverage(A, B))
}
