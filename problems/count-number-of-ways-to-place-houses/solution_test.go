package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	n int

	expected int
}

var data []testData = []testData{
	testData{
		n:        3,
		expected: 25,
	},
	testData{
		n:        888,
		expected: 16291105,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, countHousePlacements(item.n))
	}
}
