package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	nums := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}
	expected := []int{1, 4, 2, 7, 5, 3, 8, 6, 9}
	assert.Equal(t, expected, findDiagonalOrder(nums))
}

func TestSol1(t *testing.T) {
	nums := [][]int{[]int{1, 2, 3, 4, 5, 6}}
	expected := []int{1, 2, 3, 4, 5, 6}
	assert.Equal(t, expected, findDiagonalOrder(nums))
}

func TestSol2(t *testing.T) {
	nums := [][]int{[]int{1, 2, 3}, []int{4}, []int{5, 6, 7}, []int{8}, []int{9, 10, 11}}
	expected := []int{1, 4, 2, 5, 3, 8, 6, 9, 7, 10, 11}
	assert.Equal(t, expected, findDiagonalOrder(nums))
}

func TestSol3(t *testing.T) {
	nums := [][]int{[]int{1, 2, 3, 4, 5}, []int{6, 7}, []int{8}, []int{9, 10, 11}, []int{12, 13, 14, 15, 16}}
	expected := []int{1, 6, 2, 8, 7, 3, 9, 4, 12, 10, 5, 13, 11, 14, 15, 16}
	assert.Equal(t, expected, findDiagonalOrder(nums))
}

func TestSol4(t *testing.T) {
	nums := [][]int{[]int{14, 12, 19, 16, 9}, []int{13, 14, 15, 8, 11}, []int{11, 13, 1}}
	expected := []int{14, 13, 12, 11, 14, 19, 13, 15, 16, 1, 8, 9, 11}
	assert.Equal(t, expected, findDiagonalOrder(nums))
}
