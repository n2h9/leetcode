package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	nums1 := []int{1, 7, 11}
	nums2 := []int{2, 4, 6}
	k := 3
	expected := [][]int{
		[]int{1, 2},
		[]int{1, 4},
		[]int{1, 6},
	}

	assert.Equal(t, expected, kSmallestPairs(nums1, nums2, k))
}

func TestSol1(t *testing.T) {
	nums1 := []int{1, 1, 2}
	nums2 := []int{1, 2, 3}
	k := 2
	expected := [][]int{
		[]int{1, 1},
		[]int{1, 1},
	}

	assert.Equal(t, expected, kSmallestPairs(nums1, nums2, k))
}

func TestSol2(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3}
	k := 3
	expected := [][]int{
		[]int{1, 3},
		[]int{2, 3},
	}

	assert.Equal(t, expected, kSmallestPairs(nums1, nums2, k))
}

func TestSol3(t *testing.T) {
	nums1 := []int{4}
	nums2 := []int{1, 2}
	k := 3
	expected := [][]int{
		[]int{4, 1},
		[]int{4, 2},
	}

	assert.Equal(t, expected, kSmallestPairs(nums1, nums2, k))
}
