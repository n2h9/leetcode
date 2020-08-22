package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	expected := 5

	assert.Equal(t, expected, findKthLargest(nums, k))
}

func TestSol1(t *testing.T) {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4
	expected := 4

	assert.Equal(t, expected, findKthLargest(nums, k))
}
