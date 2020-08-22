package solution

type value struct {
	vertice int
	stops   int
	cost    int
}

func (v *value) Priority() int {
	// as we have max heap to turn it to min heap simply use negative values for cost
	// (^v.cost + 1) === (-1 * v.cost)
	return ^v.cost + 1
}

func newValue(vertice, stops, cost int) *value {
	return &value{
		vertice: vertice,
		stops:   stops,
		cost:    cost,
	}
}

type queue struct {
	storage *Heap
}

func newPriorityQueue() *queue {
	return &queue{
		storage: NewHeap([]HeapItem{}),
	}
}

func (q *queue) isEmpty() bool {
	return q.storage.IsEmpty()
}

func (q *queue) top() *value {
	return (q.storage.Top()).(*value)
}

func (q *queue) pop() {
	if q.isEmpty() {
		return
	}
	q.storage.Pop()
}

func (q *queue) push(v *value) {
	q.storage.Push(v)
}
