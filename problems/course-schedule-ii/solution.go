package solution

func findOrder(numCourses int, prerequisites [][]int) []int {
	if numCourses <= 0 {
		return []int{}
	}

	adjList := make([][]int, numCourses)
	inEdgeCount := make([]int, numCourses)
	for _, dep := range prerequisites {
		adjList[dep[1]] = append(adjList[dep[1]], dep[0])
		inEdgeCount[dep[0]]++
	}

	q := newQueue()
	for v, c := range inEdgeCount {
		if c == 0 {
			q.enqueue(newNode(v))
		}
	}

	resOrder := make([]int, 0, numCourses)
	for !q.isEmpty() {
		v := q.top().val
		q.dequeue()
		resOrder = append(resOrder, v)
		for _, u := range adjList[v] {
			inEdgeCount[u]--
			if inEdgeCount[u] == 0 {
				q.enqueue(newNode(u))
			}
		}
	}

	// imposible, graph has cycles
	if len(resOrder) != numCourses {
		return []int{}
	}

	return resOrder
}

type node struct {
	val  int
	next *node
}

func newNode(v int) *node {
	return &node{val: v}
}

type queue struct {
	first *node
	last  *node
}

func newQueue() *queue {
	return &queue{}
}

func (q *queue) isEmpty() bool {
	return q.first == nil
}

// queue is not empty
func (q *queue) top() *node {
	return q.first
}

// queue is not empty
func (q *queue) dequeue() {
	n := q.first
	q.first = q.first.next
	n.next = nil
	if q.first == nil {
		q.last = nil
	}
}

func (q *queue) enqueue(n *node) {
	if q.last == nil {
		q.last = n
		q.first = n
		return
	}
	q.last.next = n
	q.last = q.last.next
}
