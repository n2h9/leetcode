package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	hats := [][]int{[]int{3, 4}, []int{4, 5}, []int{5}}
	expected := 1
	assert.Equal(t, expected, numberWays(hats))
}
