package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	days := []int{1, 4, 6, 7, 8, 20}
	costs := []int{2, 7, 15}
	expected := 11

	assert.Equal(t, expected, mincostTickets(days, costs))
}

func TestSol1(t *testing.T) {
	days := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}
	costs := []int{2, 7, 15}
	expected := 17

	assert.Equal(t, expected, mincostTickets(days, costs))
}

func TestSol2(t *testing.T) {
	days := []int{1, 30}
	costs := []int{11, 3, 20}
	expected := 6

	assert.Equal(t, expected, mincostTickets(days, costs))
}
