package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	s := "leetcode"
	expected := "result"
	assert.Equal(t, expected, sol(s))
}
