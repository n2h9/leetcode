package defang

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func defangIPaddr(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}

func TestDefangIPaddr(t *testing.T) {
	address := "1.1.1.1"
	expected := "1[.]1[.]1[.]1"
	assert.Equal(t, expected, defangIPaddr(address), "Dots should be replaced with [.]")
}
