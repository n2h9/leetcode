package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	A := 1
	expected := 2

	assert.Equal(t, expected, f(A))
}
