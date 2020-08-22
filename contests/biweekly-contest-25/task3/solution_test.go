package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	s1 := "abc"
	s2 := "xya"
	expected := true
	assert.Equal(t, expected, checkIfCanBreak(s1, s2))
}

func TestSol1(t *testing.T) {
	s1 := "abe"
	s2 := "acd"
	expected := false
	assert.Equal(t, expected, checkIfCanBreak(s1, s2))
}

func TestSol2(t *testing.T) {
	s1 := "leetcodee"
	s2 := "interview"
	expected := true
	assert.Equal(t, expected, checkIfCanBreak(s1, s2))
}

func TestSol3(t *testing.T) {
	s1 := "aaaaaa"
	s2 := "aaaaaa"
	expected := true
	assert.Equal(t, expected, checkIfCanBreak(s1, s2))
}

func TestSol4(t *testing.T) {
	s1 := "aaaaaax"
	s2 := "aaaaaac"
	expected := true
	assert.Equal(t, expected, checkIfCanBreak(s1, s2))
}

func TestSol5(t *testing.T) {
	s1 := "aaaaaaxa"
	s2 := "aaaaaacb"
	expected := false
	assert.Equal(t, expected, checkIfCanBreak(s1, s2))
}
