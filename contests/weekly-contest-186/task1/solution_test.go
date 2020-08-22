package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	s := "00111"
	expected := 5
	assert.Equal(t, expected, maxScore(s))
}

func TestSol1(t *testing.T) {
	s := "1111"
	expected := 3
	assert.Equal(t, expected, maxScore(s))
}

func TestSol2(t *testing.T) {
	s := "00"
	expected := 1
	assert.Equal(t, expected, maxScore(s))
}
