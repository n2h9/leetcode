package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	s := "leetcode"
	expected := ""
	assert.Equal(t, expected, reformat(s))
}

func TestSol1(t *testing.T) {
	s := "2020202"
	expected := ""
	assert.Equal(t, expected, reformat(s))
}

func TestSol2(t *testing.T) {
	s := "202abcde"
	expected := ""
	assert.Equal(t, expected, reformat(s))
}

func TestSol3(t *testing.T) {
	s := "covid2019"
	expected := "c2o0v1i9d"
	assert.Equal(t, expected, reformat(s))
}
