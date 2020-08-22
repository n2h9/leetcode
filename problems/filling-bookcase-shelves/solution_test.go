package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	books := [][]int{
		[]int{1, 1},
		[]int{2, 3},
		[]int{2, 3},
		[]int{1, 1},
		[]int{1, 1},
		[]int{1, 1},
		[]int{1, 2},
	}
	shelfWidth := 4
	expected := 6

	assert.Equal(t, expected, minHeightShelves(books, shelfWidth))
}

func TestSol1(t *testing.T) {
	books := [][]int{
		[]int{1, 3},
		[]int{2, 4},
		[]int{3, 2},
	}
	shelfWidth := 6
	expected := 4

	assert.Equal(t, expected, minHeightShelves(books, shelfWidth))
}
