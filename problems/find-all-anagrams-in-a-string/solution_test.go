package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	s := "cbaebabacd"
	p := "abc"
	expected := []int{0, 6}
	assert.Equal(t, expected, findAnagrams(s, p))
}

func TestSol1(t *testing.T) {
	s := "abab"
	p := "ab"
	expected := []int{0, 1, 2}
	assert.Equal(t, expected, findAnagrams(s, p))
}

func TestSol2(t *testing.T) {
	s := "abbab"
	p := "abb"
	expected := []int{0, 1, 2}
	assert.Equal(t, expected, findAnagrams(s, p))
}

func TestSol3(t *testing.T) {
	s := "ababc"
	p := "abc"
	expected := []int{2}
	assert.Equal(t, expected, findAnagrams(s, p))
}

func TestSol4(t *testing.T) {
	s := ""
	p := ""
	expected := []int{}
	assert.Equal(t, expected, findAnagrams(s, p))
}
