package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	num := 555
	expected := 888
	assert.Equal(t, expected, maxDiff(num))
}

func TestSol1(t *testing.T) {
	num := 9
	expected := 8
	assert.Equal(t, expected, maxDiff(num))
}

func TestSol2(t *testing.T) {
	num := 123456
	expected := 820000
	assert.Equal(t, expected, maxDiff(num))
}

func TestSol3(t *testing.T) {
	num := 10000
	expected := 80000
	assert.Equal(t, expected, maxDiff(num))
}

func TestSol4(t *testing.T) {
	num := 9288
	expected := 8700
	assert.Equal(t, expected, maxDiff(num))
}

func TestSol5(t *testing.T) {
	num := 111
	expected := 888
	assert.Equal(t, expected, maxDiff(num))
}

func TestSol6(t *testing.T) {
	num := 101
	expected := 808
	assert.Equal(t, expected, maxDiff(num))
}
