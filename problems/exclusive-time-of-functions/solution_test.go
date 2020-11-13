package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	n    int
	logs []string

	expected []int
}

var data []testData = []testData{
	testData{
		n:        2,
		logs:     []string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"},
		expected: []int{3, 4},
	},
	testData{
		n:        1,
		logs:     []string{"0:start:0", "0:start:2", "0:end:5", "0:start:6", "0:end:6", "0:end:7"},
		expected: []int{8},
	},
	testData{
		n:        2,
		logs:     []string{"0:start:0", "0:start:2", "0:end:5", "1:start:6", "1:end:6", "0:end:7"},
		expected: []int{7, 1},
	},
	testData{
		n:        2,
		logs:     []string{"0:start:0", "0:start:2", "0:end:5", "1:start:7", "1:end:7", "0:end:8"},
		expected: []int{8, 1},
	},
	testData{
		n:        1,
		logs:     []string{"0:start:0", "0:end:0"},
		expected: []int{1},
	},
	testData{
		n:        1,
		logs:     []string{"0:start:0", "0:start:1", "0:start:2", "0:end:3", "0:end:4", "0:end:5"},
		expected: []int{6},
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, exclusiveTime(item.n, item.logs))
	}
}
