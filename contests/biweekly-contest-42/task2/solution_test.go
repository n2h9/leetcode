package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	customers [][]int

	expected float64
}

var data []testData = []testData{
	testData{
		customers:        [][]int{[]int{1,2},[]int{2,5},[]int{4,3}},
		expected: 5.0,
	},
	testData{
		customers:        [][]int{[]int{5,2},[]int{5,4},[]int{10,3}, []int{20,1}},
		expected: 3.25,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, averageWaitingTime(item.customers))
	}
}
