package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	x := 2.0
	n := 4
	expected := 16.0

	assert.Equal(t, expected, myPow(x, n))
}

func TestSol1(t *testing.T) {
	x := 2.0
	n := 5
	expected := 32.0

	assert.Equal(t, expected, myPow(x, n))
}

func TestSol2(t *testing.T) {
	x := 2.0
	n := -4
	expected := 0.0625

	assert.Equal(t, expected, myPow(x, n))
}

func TestSol4(t *testing.T) {
	x := 2.0
	n := -5
	expected := 0.03125

	assert.Equal(t, expected, myPow(x, n))
}

func TestSol5(t *testing.T) {
	x := 8.95371
	n := -1
	expected := 0.11168554710840535

	assert.Equal(t, expected, myPow(x, n))
}

func TestSol6(t *testing.T) {
	x := -1.0
	n := 2147483647
	expected := -1.0

	assert.Equal(t, expected, myPow(x, n))
}
