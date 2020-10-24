package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	nums []int
	k    int

	expected int
}

var data []testData = []testData{
	// testData{
	// 	nums: []int{1, 1, 1},
	// 	k:    2,

	// 	expected: 2,
	// },
	// testData{
	// 	nums: []int{1, 0, 1, 0, 1},
	// 	k:    2,

	// 	expected: 4,
	// },
	// testData{
	// 	nums: []int{1, -1, 0, 1, -1, 1, -1, 1, -1, 1},
	// 	k:    1,

	// 	expected: 19,
	// },
	testData{
		nums: []int{89, -287, 537, -520, 962, 686, -473, -171, -544, -826, 116, 450, -184, -305, 527, 810, 6, 834, 281, 469, -259, -831, 107, -546, -809, 800, -205, -328, -172, 764, 583, 845, 576, -303, -585, -703, -410, -84, -97, 606, 71, 312, 655, 601, 931, 411, -128, 428, -874, -924, 734, 208, 250, -478, -901, 699, 847, 722, -799, 967, -657, -892, 203, 274, -476, 104, 300, -120, 937, 684, 469, 81, -224, -256, 510, 387, -820, 995, -219, -640, -77, 634, 623, -129, -656, -665, -36, -481, 650, -746, 330, 649, -41, 826, 832, 461, 86, -139, -459, 865, 239, -96, 464, -258, -626, 173, 745, 951, -885, 768, 709, -549},
		k:    994,

		expected: 10,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, subarraySum(item.nums, item.k))
	}
}
