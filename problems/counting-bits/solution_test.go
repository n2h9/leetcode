package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	A := 2
	expected := []int{0, 1, 1}

	assert.Equal(t, expected, countBits(A))
}

func TestSol1(t *testing.T) {
	A := 5
	expected := []int{0, 1, 1, 2, 1, 2}

	assert.Equal(t, expected, countBits(A))
}
