package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	expected := 5

	assert.Equal(t, expected, maxProfit(prices))
}

func TestSol1(t *testing.T) {
	prices := []int{1, 2, 3, 4, 5}
	expected := 4

	assert.Equal(t, expected, maxProfit(prices))
}

func TestSol2(t *testing.T) {
	prices := []int{7, 6, 4, 3, 1}
	expected := 0

	assert.Equal(t, expected, maxProfit(prices))
}
