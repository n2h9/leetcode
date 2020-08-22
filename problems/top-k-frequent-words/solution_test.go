package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	words := []string{"i", "love", "leetcode", "i", "love", "coding"}
	k := 2
	expected := []string{"i", "love"}

	assert.Equal(t, expected, topKFrequent(words, k))
}

func TestSol1(t *testing.T) {
	words := []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}
	k := 4
	expected := []string{"the", "is", "sunny", "day"}

	assert.Equal(t, expected, topKFrequent(words, k))
}

func TestSol2(t *testing.T) {
	words := []string{"the", "the", "the", "the", "the", "the", "the", "the", "the", "the"}
	k := 4
	expected := []string{"the"}

	assert.Equal(t, expected, topKFrequent(words, k))
}

func TestSol3(t *testing.T) {
	words := []string{}
	k := 4
	expected := []string{}

	assert.Equal(t, expected, topKFrequent(words, k))
}

func TestSol4(t *testing.T) {
	words := []string{"b", "b", "a", "a"}
	k := 4
	expected := []string{"a", "b"}

	assert.Equal(t, expected, topKFrequent(words, k))
}

func TestSol5(t *testing.T) {
	words := []string{"b", "b", "b", "a", "a"}
	k := 4
	expected := []string{"b", "a"}

	assert.Equal(t, expected, topKFrequent(words, k))
}
