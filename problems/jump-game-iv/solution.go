package solution

func minJumps(arr []int) int {
	const MAX_INT = 1 << 31

	n := len(arr)

	dp := make([]int, n)
	valMap := make(map[int][]int, n)

	for i := 0; i < n; i++ {
		valMap[arr[i]] = append(valMap[arr[i]], i)
		dp[i] = MAX_INT
	}
	dp[0] = 0

	q := newQueue()
	q.push(newNode(0))

	for !q.isEmpty() {
		i := q.top().value
		q.pop()
		// next jump len from 0
		jv := dp[i]+1
		// indexes possible to jump from i
		jixs := []int{}
		if i-1 >= 0 {
			jixs = append(jixs, i-1)
		}
		if i+1 < n {
			jixs = append(jixs, i+1)
		}
		jixs = append(jixs, valMap[arr[i]]...)
		for _, j := range jixs {
			if jv < dp[j] {
				dp[j] = jv
				q.push(newNode(j))
			}
		}
		valMap[arr[i]] = nil
	}
	
	return dp[n-1]
}

type node struct {
	value int
	next *node
}

func newNode(value int) *node {
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

	q.tail.next = n
	q.tail = n
}