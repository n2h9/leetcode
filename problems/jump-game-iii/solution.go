package solution

func canReach(arr []int, start int) bool {
	q := newQueue()
	q.push(newNode(start))
	visited := make([]bool, len(arr))
	for !q.isEmpty() {
		i := q.top().val
		q.pop()
		if arr[i] == 0 {
			return true
		}
		for _, move := range []int{i + arr[i], i - arr[i]} {
			if move >= 0 && move < len(arr) && !visited[move] {
				q.push(newNode(move))
				visited[move] = true
			}
		}
	}
	return false
}

type node struct {
	next *node
	val  int
}

func newNode(val int) *node {
	return &node{val: val}
}

type queue struct {
	tail *node
	head *node
}

func newQueue() *queue {
	return &queue{}
}

func (q *queue) isEmpty() bool {
	return q.head == nil
}

func (q *queue) top() *node {
	return q.head
}

func (q *queue) pop() {
	n := q.head
	q.head = q.head.next
	n.next = nil
	if q.head == nil {
		q.tail = nil
	}
}

func (q *queue) push(n *node) {
	if q.head == nil {
		q.head = n
		q.tail = n
		return
	}

	n.next = q.head
	q.head = n
}
