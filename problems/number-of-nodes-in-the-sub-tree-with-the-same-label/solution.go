package solution

func countSubTrees(n int, edges [][]int, labels string) []int {
	parents := make([]int, n)
	adjacencyList := make([][]int, n)

	for i := 0; i < n; i++ {
		parents[i] = -1
		// 3 but not 2 because there is no gourantie that edge[0] is parent and edge[1] is child
		// so add one as a child of the second for both nodes
		// so max amount of adjacent nodes per node is 3
		adjacencyList[i] = make([]int, 0, 3)
	}

	for _, edge := range edges {
		adjacencyList[edge[0]] = append(adjacencyList[edge[0]], edge[1])
		adjacencyList[edge[1]] = append(adjacencyList[edge[1]], edge[0])
	}

	q := newQueue()
	s := newStack()

	q.push(0)

	for !q.isEmpty() {
		node := q.pop()
		s.push(node)

		for _, neighbour := range adjacencyList[node] {
			if neighbour != parents[node] {
				parents[neighbour] = node
				q.push(neighbour)
			}
		}
	}

	labelsCount := make([]map[byte]int, n)
	for i := 0; i < n; i++ {
		labelsCount[i] = make(map[byte]int, 26)
	}

	for !s.isEmpty() {
		node := s.pop()
		labelsCount[node][labels[node]]++ // increment current node labels counter
		if parents[node] == -1 {
			continue
		}
		// add own labels counter to parent
		for k, v := range labelsCount[node] {
			labelsCount[parents[node]][k] += v
		}
	}

	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = labelsCount[i][labels[i]]
	}

	return res
}

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

// ----
// stack
type stack []int

func newStack() stack {
	return make([]int, 0, 32)
}

func (s *stack) isEmpty() bool {
	return len((*s)) <= 0
}

func (s *stack) push(node int) {
	(*s) = append((*s), node)
}

// panic on empty stack
func (s *stack) pop() int {
	v := (*s)[len((*s))-1]
	(*s) = (*s)[:len((*s))-1]
	return v
}

// end of stack
// ----
