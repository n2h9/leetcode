package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	students []int
	sandwiches []int

	expected int
}

var data []testData = []testData{
	testData{
		students:        []int{1,1,0,0},
		sandwiches:		[]int{0,1,0,1},
		expected: 0,
	},
	testData{
		students:        []int{1,1,1,0,0,1},
		sandwiches:		[]int{1,0,0,0,1,1},
		expected: 3,
	},
	testData{
		students:        []int{1},
		sandwiches:		[]int{0},
		expected: 1,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, countStudents(item.students, item.sandwiches))
	}
}
