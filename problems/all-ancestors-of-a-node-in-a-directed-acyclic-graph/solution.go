package solution

func getAncestors(n int, edges [][]int) [][]int {
	merge := func(arrays ...[]int) []int {
		m := make(map[int]bool)
		for _, arr := range arrays {
			for _, val := range arr {
				m[val] = true
			}
		}

		res := make([]int, 0, n)
		for i := 0; i < n; i++ {
			if m[i] {
				res = append(res, i)
			}
		}

		return res
	}

	children := make([][]int, n)
	parents := make([][]int, n)

	for i := 0; i < n; i++ {
		children[i] = make([]int, 0, n-1)
		parents[i] = make([]int, 0, n-1)
	}

	for _, edge := range edges {
		children[edge[0]] = append(children[edge[0]], edge[1])
		parents[edge[1]] = append(parents[edge[1]], edge[0])
	}

	q := newQueue()

	// add nodes with no parents to queue
	for i := 0; i < n; i++ {
		if len(parents[i]) <= 0 {
			q.push(i)
		}
	}

	ancestors := make([][]int, n)

	isProcessed := make(map[int]bool, n)
	for !q.isEmpty() {
		node := q.pop()
		if isProcessed[node] {
			continue
		}

		// not yet all parents processed, skip this node, it will appear later in a queue when all parents will be processed
		notAllParentsProcessedYet := false
		for _, parent := range parents[node] {
			if !isProcessed[parent] {
				notAllParentsProcessedYet = true
				break
			}
		}
		if notAllParentsProcessedYet {
			continue
		}

		parentsAncestors := make([][]int, 0, len(parents[node]))
		for _, parent := range parents[node] {
			parentsAncestors = append(parentsAncestors, ancestors[parent])
		}

		ancestors[node] = merge(
			append(
				parentsAncestors,
				parents[node],
			)...,
		)

		for _, child := range children[node] {
			q.push(child)
		}
		isProcessed[node] = true
	}

	return ancestors
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
