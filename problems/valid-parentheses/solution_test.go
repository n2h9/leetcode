package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	s := "()"
	expected := true

	assert.Equal(t, expected, isValid(s))
}

func TestSol1(t *testing.T) {
	s := "()[]{}"
	expected := true

	assert.Equal(t, expected, isValid(s))
}

func TestSol2(t *testing.T) {
	s := "(]"
	expected := false

	assert.Equal(t, expected, isValid(s))
}

func TestSol3(t *testing.T) {
	s := "([)]"
	expected := false

	assert.Equal(t, expected, isValid(s))
}

func TestSol4(t *testing.T) {
	s := "{[]}"
	expected := true

	assert.Equal(t, expected, isValid(s))
}

func TestSol5(t *testing.T) {
	s := "{[(("
	expected := false

	assert.Equal(t, expected, isValid(s))
}

func TestSol6(t *testing.T) {
	s := ")()"
	expected := false

	assert.Equal(t, expected, isValid(s))
}
