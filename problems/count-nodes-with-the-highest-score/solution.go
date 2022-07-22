package solution

func countHighestScoreNodes(parents []int) int {
	n := len(parents)

	// [root] -> [left, right]
	// contains children of evety node
	children := make([][]int, n)
	for i := 1; i < n; i++ {
		if children[parents[i]] == nil {
			children[parents[i]] = make([]int, 0, 2)
		}
		children[parents[i]] = append(children[parents[i]], i)
	}

	q := newQueue()
	s := newStack()

	q.push(0) // push root node to queue

	// go through each level of the tree
	// adding each node to the stack
	// so that when stack will be processed, all nodes will be processed
	// in order: level by level, the deepest leaves first
	for !q.isEmpty() {
		node := q.pop()
		s.push(node)

		for _, child := range children[node] {
			q.push(child)
		}
	}

	// contains the size of node i
	sizes := make([]int, n)

	for !s.isEmpty() {
		node := s.pop()

		// count node itself
		sizes[node]++

		// add a size of a child to a size of a parent
		if parents[node] != -1 {
			sizes[parents[node]] += sizes[node]
		}
	}

	maxScore := 0
	scoreNum := make(map[int]int)

	// go through all nodes except root
	for i := 0; i < n; i++ {
		// score = size[left] * size[right] * (n - size[i])
		score := 1
		for _, child := range children[i] {
			score *= sizes[child]
		}
		if i > 0 { // means contains parent
			score *= (n - sizes[i])
		}

		scoreNum[score]++
		if score > maxScore {
			maxScore = score
		}
	}

	return scoreNum[maxScore]
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
