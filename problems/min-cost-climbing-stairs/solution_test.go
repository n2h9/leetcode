package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	cost := []int{10, 15, 20}
	expected := 15

	assert.Equal(t, expected, minCostClimbingStairs(cost))
}

func TestSol1(t *testing.T) {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	expected := 6

	assert.Equal(t, expected, minCostClimbingStairs(cost))
}
