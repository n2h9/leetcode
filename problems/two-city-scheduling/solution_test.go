package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	costs := [][]int{
		[]int{30, 200},
		[]int{400, 50},
		[]int{30, 20},
		[]int{10, 20},
	}
	expected := 110
	assert.Equal(t, expected, twoCitySchedCost(costs))
}

func TestSol1(t *testing.T) {
	costs := [][]int{
		[]int{110, 100},
		[]int{210, 200},
	}
	expected := 310
	assert.Equal(t, expected, twoCitySchedCost(costs))
}

func TestSol2(t *testing.T) {
	costs := [][]int{
		[]int{200, 100},
		[]int{200, 100},
	}
	expected := 300
	assert.Equal(t, expected, twoCitySchedCost(costs))
}
