package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	expected := []int{5, 6, 7, 1, 2, 3, 4}

	rotate(nums, k)
	assert.Equal(t, expected, nums)
}

func TestSol1(t *testing.T) {
	nums := []int{-1, -100, 3, 99}
	k := 2
	expected := []int{3, 99, -1, -100}

	rotate(nums, k)
	assert.Equal(t, expected, nums)
}

func TestSol2(t *testing.T) {
	nums := []int{1, 2, 3, 999, 5000, -80000}
	k := 100
	expected := []int{3, 999, 5000, -80000, 1, 2}

	rotate(nums, k)
	assert.Equal(t, expected, nums)
}
