package solution

func ladderLength(beginWord string, endWord string, wordList []string) int {
	n := len(wordList)
	if n <= 0 {
		return 0
	}
	noEndWord := true
	al := make([][]int, n)
	for i := 0; i < n; i++ {
		al[i] = make([]int, 0, n)
		if endWord == wordList[i] {
			noEndWord = false
		}
	}
	if noEndWord {
		return 0
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isOneCharDist(wordList[i], wordList[j]) {
				al[i] = append(al[i], j)
				al[j] = append(al[j], i)
			}
		}
	}

	visited := make([]bool, n)
	q := newQueue()
	for i := 0; i < n; i++ {
		if isOneCharDist(beginWord, wordList[i]) {
			if wordList[i] == endWord {
				return 2
			}
			q.push(newNode(&item{vertex: i, len: 2}))
		}
	}

	for !q.isEmpty() {
		vertex, nextLen := q.top().value.vertex, q.top().value.len+1
		q.pop()
		if visited[vertex] {
			continue
		}
		visited[vertex] = true
		for _, w := range al[vertex] {
			if wordList[w] == endWord {
				return nextLen
			}
			if !visited[w] {
				q.push(newNode(&item{vertex: w, len: nextLen}))
			}
		}
	}

	return 0
}

func isOneCharDist(a, b string) bool {
	diffCount := 0
	n := len(a)
	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			diffCount++
			if diffCount > 1 {
				return false
			}
		}
	}
	return true
}

type item struct {
	vertex int
	len    int
}

type node struct {
	value *item
	next  *node
}

func newNode(value *item) *node {
	return &node{value: value}
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
	h := q.head
	q.head = q.head.next
	h.next = nil

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

	q.tail.next = n
	q.tail = n
}
