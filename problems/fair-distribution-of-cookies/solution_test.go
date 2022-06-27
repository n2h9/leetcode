package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	cookies []int
	k       int

	expected int
}

var data []testData = []testData{
	{
		cookies:  []int{8, 15, 10, 20, 8},
		k:        2,
		expected: 31,
	},
	{
		cookies:  []int{6, 1, 3, 2, 2, 4, 1, 2},
		k:        3,
		expected: 7,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, distributeCookies(item.cookies, item.k))
	}
}
