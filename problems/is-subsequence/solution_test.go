package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	s := "abc"
	s1 := "ahbgdc"
	expected := true
	assert.Equal(t, expected, isSubsequence(s, s1))
}

func TestSol1(t *testing.T) {
	s := "axc"
	s1 := "ahbgdc"
	expected := false
	assert.Equal(t, expected, isSubsequence(s, s1))
}
