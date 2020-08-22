package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	words := []string{"hello", "leetcode"}
	order := "hlabcdefgijkmnopqrstuvwxyz"

	expected := true

	assert.Equal(t, expected, isAlienSorted(words, order))
}

func TestSol1(t *testing.T) {
	words := []string{"word", "world", "row"}
	order := "worldabcefghijkmnpqstuvxyz"

	expected := false

	assert.Equal(t, expected, isAlienSorted(words, order))
}

func TestSol2(t *testing.T) {
	words := []string{"apple", "app"}
	order := "abcdefghijklmnopqrstuvwxyz"

	expected := false

	assert.Equal(t, expected, isAlienSorted(words, order))
}

func TestSol3(t *testing.T) {
	words := []string{"app", "app"}
	order := "abcdefghijklmnopqrstuvwxyz"

	expected := true

	assert.Equal(t, expected, isAlienSorted(words, order))
}

func TestSol4(t *testing.T) {
	words := []string{"app", "appa"}
	order := "abcdefghijklmnopqrstuvwxyz"

	expected := true

	assert.Equal(t, expected, isAlienSorted(words, order))
}
