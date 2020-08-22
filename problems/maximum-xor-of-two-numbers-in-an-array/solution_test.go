package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	nums := []int{3, 10, 5, 25, 2, 8}
	expected := 28

	assert.Equal(t, expected, findMaximumXOR(nums))
}

func TestSol1(t *testing.T) {
	nums := []int{4, 11, 6, 26, 3, 9}
	expected := 30

	assert.Equal(t, expected, findMaximumXOR(nums))
}

func TestSol2(t *testing.T) {
	nums := []int{}
	expected := 0

	assert.Equal(t, expected, findMaximumXOR(nums))
}

func TestSol3(t *testing.T) {
	nums := []int{123}
	expected := 0

	assert.Equal(t, expected, findMaximumXOR(nums))
}

func TestSol4(t *testing.T) {
	nums := []int{123, 0}
	expected := 123

	assert.Equal(t, expected, findMaximumXOR(nums))
}

func TestSol5(t *testing.T) {
	nums := []int{9, 6}
	expected := 15

	assert.Equal(t, expected, findMaximumXOR(nums))
}
