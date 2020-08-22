package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	a := "11"
	b := "1"

	expected := "100"

	assert.Equal(t, expected, addBinary(a, b))
}

func TestSol1(t *testing.T) {
	a := "1010"
	b := "1011"

	expected := "10101"

	assert.Equal(t, expected, addBinary(a, b))
}

func TestSol2(t *testing.T) {
	a := "0"
	b := "1111"

	expected := "1111"

	assert.Equal(t, expected, addBinary(a, b))
}

func TestSol3(t *testing.T) {
	a := "0"
	b := "0"

	expected := "0"

	assert.Equal(t, expected, addBinary(a, b))
}

func TestSol4(t *testing.T) {
	a := "11111"
	b := "1"

	expected := "100000"

	assert.Equal(t, expected, addBinary(a, b))
}

func TestSol5(t *testing.T) {
	a := "1111"
	b := "1111"

	expected := "11110"

	assert.Equal(t, expected, addBinary(a, b))
}
