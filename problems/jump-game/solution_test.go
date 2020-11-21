package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	nums []int

	expected bool
}

var data []testData = []testData{
	testData{
		nums:        []int{2,3,1,1,4},
		expected: true,
	},
	testData{
		nums:        []int{3,2,1,0,4},
		expected: false,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, canJump(item.nums))
	}
}
