package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	expected := []int{24, 12, 8, 6}

	assert.Equal(t, expected, productExceptSelf(nums))
}

func TestSol1(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	expected := []int{720, 360, 240, 180, 144, 120}

	assert.Equal(t, expected, productExceptSelf(nums))
}

func TestSol2(t *testing.T) {
	nums := []int{100, 200}
	expected := []int{200, 100}

	assert.Equal(t, expected, productExceptSelf(nums))
}

func TestSol3(t *testing.T) {
	nums := []int{10, 20, 30}
	expected := []int{600, 300, 200}

	assert.Equal(t, expected, productExceptSelf(nums))
}
