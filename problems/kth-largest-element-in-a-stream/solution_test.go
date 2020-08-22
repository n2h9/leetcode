package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	A := 3
	B := []int{4, 5, 8, 2}

	kthLargest := Constructor(A, B)
	assert.NotNil(t, kthLargest)
	assert.Equal(t, 4, kthLargest.Add(3))
	assert.Equal(t, 5, kthLargest.Add(5))
	assert.Equal(t, 5, kthLargest.Add(10))
	assert.Equal(t, 8, kthLargest.Add(9))
	assert.Equal(t, 8, kthLargest.Add(4))
}
