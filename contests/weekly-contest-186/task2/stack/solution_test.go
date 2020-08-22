package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	cardPoints := []int{1, 2, 3, 4, 5, 6, 1}
	k := 3
	expected := 12
	assert.Equal(t, expected, maxScore(cardPoints, k))
}

func TestSol1(t *testing.T) {
	cardPoints := []int{2, 2, 2}
	k := 2
	expected := 4
	assert.Equal(t, expected, maxScore(cardPoints, k))
}

func TestSol2(t *testing.T) {
	cardPoints := []int{9, 7, 7, 9, 7, 7, 9}
	k := 7
	expected := 55
	assert.Equal(t, expected, maxScore(cardPoints, k))
}

func TestSol3(t *testing.T) {
	cardPoints := []int{1, 1000, 1}
	k := 1
	expected := 1
	assert.Equal(t, expected, maxScore(cardPoints, k))
}

func TestSol4(t *testing.T) {
	cardPoints := []int{1, 79, 80, 1, 1, 1, 200, 1}
	k := 3
	expected := 202
	assert.Equal(t, expected, maxScore(cardPoints, k))
}

func TestSol5(t *testing.T) {
	cardPoints := []int{30, 88, 33, 37, 18, 77, 54, 73, 31, 88, 93, 25, 18, 31, 71, 8, 97, 20, 98, 16, 65, 40, 18, 25, 13, 51, 59}
	k := 26
	expected := 1269
	assert.Equal(t, expected, maxScore(cardPoints, k))
}
