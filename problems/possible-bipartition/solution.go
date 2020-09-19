package solution

const COLOR_RED byte = 1
const COLOR_BLUE byte = 2

func possibleBipartition(N int, dislikes [][]int) bool {
	al := make([][]int, N+1)
	for _, dislike := range dislikes {
		u := dislike[0]
		v := dislike[1]
		al[u] = append(al[u], v)
		al[v] = append(al[v], u)
	}
	visited := make([]bool, N+1)
	colors := make([]byte, N+1)
	q := newQueue()
	for i := 1; i <= N; i++ {
		if visited[i] {
			continue
		}
		colors[i] = COLOR_RED
		q.push(newNode(i))
		for !q.isEmpty() {
			v := q.top().val
			q.pop()
			if visited[v] {
				continue
			}
			visited[v] = true
			neightborColor := changeColor(colors[v])
			for _, u := range al[v] {
				if colors[u] == colors[v] {
					return false
				} else if colors[u] == 0 {
					colors[u] = neightborColor
					q.push(newNode(u))
				}
			}
		}
	}
	return true
}

func changeColor(c byte) byte {
	if c == COLOR_BLUE {
		return COLOR_RED
	}
	return COLOR_BLUE
}

type node struct {
	val  int
	next *node
}

func newNode(val int) *node {
	return &node{val: val}
}

type queue struct {
	head *node
	tail *node
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
	if q.tail == nil {
		q.tail = n
		q.head = n
		return
	}
	q.tail.next = n
	q.tail = n
}
