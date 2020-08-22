package solution

type color bool

const BLUE color = false
const RED color = true

func shortestAlternatingPaths(n int, red_edges [][]int, blue_edges [][]int) []int {
	if n <= 0 {
		return []int{}
	}

	adjList := make(map[color][][]int, 2)
	adjList[RED] = make([][]int, n)
	for _, edge := range red_edges {
		adjList[RED][edge[0]] = append(adjList[RED][edge[0]], edge[1])
	}

	adjList[BLUE] = make([][]int, n)
	for _, edge := range blue_edges {
		adjList[BLUE][edge[0]] = append(adjList[BLUE][edge[0]], edge[1])
	}

	visited := make(map[color][]bool)
	minLen := make(map[color][]int)
	for _, c := range []color{BLUE, RED} {
		visited[c] = make([]bool, n)
		minLen[c] = make([]int, n)
	}
	q := NewQueue()
	q.enqueue(&queueItem{v: 0, c: BLUE, priority: 0})
	q.enqueue(&queueItem{v: 0, c: RED, priority: 0})
	visited[BLUE][0] = true
	visited[RED][0] = true

	for !q.isEmpty() {
		item := q.top()
		q.dequeue()
		nextColor := changeColor(item.c)
		for _, u := range adjList[nextColor][item.v] {
			if !visited[nextColor][u] {
				if l := minLen[item.c][item.v] + 1; minLen[nextColor][u] == 0 || l < minLen[nextColor][u] {
					minLen[nextColor][u] = l
				}
				q.enqueue(
					&queueItem{
						v: u,
						c: nextColor,
						// negative value for min priority queue
						priority: ^minLen[nextColor][u] + 1,
					},
				)
				visited[nextColor][u] = true
			}
		}
	}

	minOrNegative := func(x, y int) int {
		if x == 0 {
			if y == 0 {
				return -1
			}
			return y
		}
		if y == 0 {
			if x == 0 {
				return -1
			}
			return x
		}
		if x < y {
			return x
		}
		return y
	}

	result := make([]int, n)
	result[0] = 0
	for i := 1; i < n; i++ {
		result[i] = minOrNegative(
			minLen[BLUE][i],
			minLen[RED][i],
		)
	}

	return result
}

func changeColor(c color) color {
	if c == BLUE {
		return RED
	}
	return BLUE
}

type queueItem struct {
	v        int
	c        color
	priority int
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
			value:    item,
			priority: item.priority,
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
