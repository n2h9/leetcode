package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	prices := []int{8, 4, 6, 2, 3}
	expected := []int{4, 2, 4, 2, 3}

	assert.Equal(t, expected, finalPrices(prices))
}

func TestSol1(t *testing.T) {
	prices := []int{1, 2, 3, 4, 5}
	expected := []int{1, 2, 3, 4, 5}

	assert.Equal(t, expected, finalPrices(prices))
}

func TestSol2(t *testing.T) {
	prices := []int{10, 1, 1, 6}
	expected := []int{9, 0, 1, 6}

	assert.Equal(t, expected, finalPrices(prices))
}

func TestSol3(t *testing.T) {
	prices := []int{1, 1, 1, 1}
	expected := []int{0, 0, 0, 1}

	assert.Equal(t, expected, finalPrices(prices))
}
