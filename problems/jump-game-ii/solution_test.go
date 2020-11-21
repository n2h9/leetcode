package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	nums []int

	expected int
}

var data []testData = []testData{
	testData{
		nums:        []int{2,3,1,1,4},
		expected: 2,
	},
	testData{
		nums:        []int{2,3,0,1,4},
		expected: 2,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, jump(item.nums))
	}
}
