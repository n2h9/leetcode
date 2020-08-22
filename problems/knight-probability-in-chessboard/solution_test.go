package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	N := 3
	K := 2
	r := 0
	c := 0
	expected := 0.0625

	assert.Equal(t, expected, knightProbability(N, K, r, c))
}

func TestSol1(t *testing.T) {
	N := 3
	K := 3
	r := 0
	c := 0
	expected := 0.015625

	assert.Equal(t, expected, knightProbability(N, K, r, c))
}

func TestSol2(t *testing.T) {
	N := 4
	K := 2
	r := 2
	c := 1
	expected := 0.15625

	assert.Equal(t, expected, knightProbability(N, K, r, c))
}

func TestSol3(t *testing.T) {
	N := 3
	K := 1
	r := 1
	c := 1
	expected := 0.0

	assert.Equal(t, expected, knightProbability(N, K, r, c))
}

func TestSol4(t *testing.T) {
	N := 3
	K := 3
	r := 1
	c := 1
	expected := 0.0

	assert.Equal(t, expected, knightProbability(N, K, r, c))
}

// func TestSol5(t *testing.T) {
// 	N := 8
// 	K := 30
// 	r := 6
// 	c := 4
// 	expected := 0.00019052566298333645

// 	assert.Equal(t, expected, knightProbability(N, K, r, c))
// }
