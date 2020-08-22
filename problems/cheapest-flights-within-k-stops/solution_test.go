package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSol0(t *testing.T) {
	A := 3
	B := [][]int{
		[]int{0, 1, 100},
		[]int{1, 2, 100},
		[]int{0, 2, 500},
	}
	C := 0
	D := 2
	E := 1
	expected := 200
	assert.Equal(t, expected, findCheapestPrice(A, B, C, D, E))
}

func TestSol1(t *testing.T) {
	A := 3
	B := [][]int{
		[]int{0, 1, 100},
		[]int{1, 2, 100},
		[]int{0, 2, 500},
	}
	C := 0
	D := 2
	E := 0
	expected := 500
	assert.Equal(t, expected, findCheapestPrice(A, B, C, D, E))
}

func TestSol2(t *testing.T) {
	A := 4
	B := [][]int{
		[]int{0, 1, 1},
		[]int{0, 2, 5},
		[]int{1, 2, 1},
		[]int{2, 3, 1},
	}
	C := 0
	D := 3
	E := 1
	expected := 6
	assert.Equal(t, expected, findCheapestPrice(A, B, C, D, E))
}

func TestSol3(t *testing.T) {
	A := 5
	B := [][]int{
		[]int{0, 1, 100},
		[]int{0, 2, 500},
		[]int{1, 2, 100},
		[]int{2, 3, 100},
		[]int{3, 4, 100},
	}
	C := 0
	D := 4
	E := 2
	expected := 700
	assert.Equal(t, expected, findCheapestPrice(A, B, C, D, E))
}

func TestSol4(t *testing.T) {
	A := 6
	B := [][]int{
		[]int{0, 1, 100},
		[]int{0, 2, 500},
		[]int{1, 2, 100},
		[]int{2, 3, 100},
		[]int{3, 4, 100},
		[]int{4, 5, 100},
	}
	C := 0
	D := 5
	E := 2
	expected := -1
	assert.Equal(t, expected, findCheapestPrice(A, B, C, D, E))
}

func TestSol5(t *testing.T) {
	A := 6
	B := [][]int{
		[]int{0, 1, 100},
		[]int{0, 2, 500},
		[]int{1, 2, 100},
		[]int{2, 3, 100},
		[]int{3, 4, 100},
		[]int{4, 5, 100},
	}
	C := 0
	D := 5
	E := 5
	expected := 500
	assert.Equal(t, expected, findCheapestPrice(A, B, C, D, E))
}

func TestSol6(t *testing.T) {
	q := newPriorityQueue()
	q.push(newValue(0, 0, 100))
	q.push(newValue(0, 0, 101))
	q.push(newValue(0, 0, 99))
	q.push(newValue(0, 0, 200))
	q.push(newValue(0, 0, 100))

	assert.Equal(t, 99, q.top().cost)
	q.pop()
	assert.Equal(t, 100, q.top().cost)
	q.pop()
	assert.Equal(t, 100, q.top().cost)
	q.pop()
	assert.Equal(t, 101, q.top().cost)
	q.pop()
	assert.Equal(t, 200, q.top().cost)
	q.pop()
	assert.Equal(t, true, q.isEmpty())
}
