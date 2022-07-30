package solution

// ----
// queue

type queue []int

func newQueue() queue {
	return make([]int, 0, 32)
}

func (q *queue) isEmpty() bool {
	return len(*q) <= 0
}

func (q *queue) push(node int) {
	(*q) = append(*q, node)
}

// panic on empty queue
func (q *queue) pop() int {
	v := (*q)[0]
	(*q) = (*q)[1:]
	return v
}

// end of queue
// ----
