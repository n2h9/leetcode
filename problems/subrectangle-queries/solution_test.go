package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	subrectangleQueries := Constructor(
		[][]int{
			[]int{1, 2, 1},
			[]int{4, 3, 4},
			[]int{3, 2, 1},
			[]int{1, 1, 1},
		},
	)
	assert.Equal(t, 1, subrectangleQueries.GetValue(0, 2))

	subrectangleQueries.UpdateSubrectangle(0, 0, 3, 2, 5)
	assert.Equal(t, 5, subrectangleQueries.GetValue(0, 2))
	assert.Equal(t, 5, subrectangleQueries.GetValue(3, 1))

	subrectangleQueries.UpdateSubrectangle(3, 0, 3, 2, 10)
	assert.Equal(t, 5, subrectangleQueries.GetValue(0, 2))
	assert.Equal(t, 10, subrectangleQueries.GetValue(3, 1))

}
