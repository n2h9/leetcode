package solution

func swimInWater(grid [][]int) int {
	n := len(grid)
	q := NewQueue()
	q.enqueue(
		&queueItem{
			i:         0,
			j:         0,
			elevation: grid[0][0],
			minDepth:  grid[0][0],
		},
	)

	visitedMap := make(map[int]map[int]bool)
	for i := 0; i < n; i++ {
		visitedMap[i] = make(map[int]bool)
	}
	visitedMap[0][0] = true

	minDepth := -1
	for !q.isEmpty() {
		cell := q.top()
		i, j, depth := cell.i, cell.j, cell.minDepth
		if i == n-1 && j == n-1 {
			minDepth = max(depth, grid[i][j])
			break
		}
		q.dequeue()
		dirs := [][]int{
			[]int{i - 1, j},
			[]int{i + 1, j},
			[]int{i, j - 1},
			[]int{i, j + 1},
		}
		for _, dir := range dirs {
			if checkIndexes(dir[0], dir[1], n) && !visitedMap[dir[0]][dir[1]] {
				visitedMap[dir[0]][dir[1]] = true
				elevation := grid[dir[0]][dir[1]]
				q.enqueue(
					&queueItem{
						i:         dir[0],
						j:         dir[1],
						elevation: elevation,
						minDepth:  max(depth, elevation),
					},
				)
			}
		}
	}

	return minDepth
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func checkIndexes(i, j, n int) bool {
	return i >= 0 && j >= 0 && i < n && j < n
}

type queueItem struct {
	i         int
	j         int
	elevation int
	minDepth  int
}

type priorityQueue struct {
	storage *heap
}

func NewQueue() *priorityQueue {
	return &priorityQueue{storage: newHeap()}
}

func (q *priorityQueue) isEmpty() bool {
	return q.storage.isEmpty()
}

func (q *priorityQueue) top() *queueItem {
	return q.storage.top().value
}

func (q *priorityQueue) enqueue(item *queueItem) {
	q.storage.push(
		&heapItem{
			value: item,
			// take opposite val of elevation to emulate min heap
			priority: ^item.elevation + 1,
		},
	)
}

func (q *priorityQueue) dequeue() {
	q.storage.pop()
}

// ---

type heapItem struct {
	value    *queueItem
	priority int
}

type heap struct {
	storage []*heapItem
}

func newHeap() *heap {
	return &heap{
		storage: make([]*heapItem, 0),
	}
}

func (h *heap) isEmpty() bool {
	return len(h.storage) <= 0
}

func (h *heap) push(item *heapItem) {
	h.storage = append(h.storage, item)
	ind := len(h.storage) - 1
	// parent index
	pi := (ind - 1) / 2
	// move item up until it is > then parent
	for ind > 0 && h.storage[ind].priority > h.storage[pi].priority {
		h.storage[ind], h.storage[pi] = h.storage[pi], h.storage[ind]
		ind = pi
		pi = (ind - 1) / 2
	}
}

func (h *heap) top() *heapItem {
	return h.storage[0]
}

func (h *heap) pop() {
	h.storage[0], h.storage[len(h.storage)-1] = h.storage[len(h.storage)-1], h.storage[0]
	h.storage = h.storage[:len(h.storage)-1]
	h._heapify(0)
}

func (h *heap) _heapify(ind int) {
	// left index, right index, max index
	li, ri, mi := 2*ind+1, 2*ind+2, ind

	for _, i := range []int{li, ri} {
		if i < len(h.storage) && h.storage[i].priority > h.storage[mi].priority {
			mi = i
		}
	}

	if mi != ind {
		h.storage[mi], h.storage[ind] = h.storage[ind], h.storage[mi]
		h._heapify(mi)
	}
}
