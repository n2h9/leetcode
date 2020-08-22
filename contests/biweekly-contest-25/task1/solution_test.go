package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	expected := []bool{true, true, true, false, true}
	assert.Equal(t, expected, kidsWithCandies(candies, extraCandies))
}

func TestSol1(t *testing.T) {
	candies := []int{12, 1, 12}
	extraCandies := 10
	expected := []bool{true, false, true}
	assert.Equal(t, expected, kidsWithCandies(candies, extraCandies))
}

func TestSol2(t *testing.T) {
	candies := []int{4, 2, 1, 1, 2}
	extraCandies := 1
	expected := []bool{true, false, false, false, false}
	assert.Equal(t, expected, kidsWithCandies(candies, extraCandies))
}
