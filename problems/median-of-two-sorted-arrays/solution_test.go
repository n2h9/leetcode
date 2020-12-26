package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	nums1 []int
	nums2 []int

	expected float64
}

var data []testData = []testData{
	testData{
		nums1:        []int{1,3,5,7,9},
		nums2:        []int{2,4,6,8,10},
		expected: 5.5,
	},
	testData{
		nums1:        []int{1,3},
		nums2:        []int{2},
		expected: 2.0,
	},
	testData{
		nums1:        []int{1,2},
		nums2:        []int{3,4},
		expected: 2.5,
	},
	testData{
		nums1:        []int{0,0},
		nums2:        []int{0,0},
		expected: 0.0,
	},
	testData{
		nums1:        []int{},
		nums2:        []int{1},
		expected: 1.0,
	},
	testData{
		nums1:        []int{2},
		nums2:        []int{},
		expected: 2.0,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, findMedianSortedArrays(item.nums1, item.nums2))
	}
}
