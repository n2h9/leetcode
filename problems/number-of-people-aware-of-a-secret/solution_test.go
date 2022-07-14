package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	n      int
	delay  int
	forget int

	expected int
}

var data []testData = []testData{
	testData{
		n:      32,
		delay:  1,
		forget: 32,

		expected: 147483634,
	},

	testData{
		n:      64,
		delay:  1,
		forget: 64,

		expected: 291172004,
	},

	testData{
		n:      128,
		delay:  1,
		forget: 128,

		expected: 639816142,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, peopleAwareOfSecret(item.n, item.delay, item.forget))
	}
}
