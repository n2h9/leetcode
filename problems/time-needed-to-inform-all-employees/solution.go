package solution

func numOfMinutes(n int, headID int, managers []int, informTime []int) int {
	children := make([][]int, n)
	for i := 0; i < n; i++ {
		children[i] = make([]int, 0)
	}
	for e, m := range managers {
		if m == -1 {
			continue
		}
		children[m] = append(children[m], e)
	}

	q := newQueue()
	q.push(headID)

	totalInformTime := make([]int, n)
	for !q.isEmpty() {
		node := q.pop()
		if len(children[node]) <= 0 {
			continue
		}

		totalInformTime[node] += informTime[node]
		if managers[node] != -1 {
			totalInformTime[node] += totalInformTime[managers[node]]
		}
		for _, child := range children[node] {
			q.push(child)
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		if totalInformTime[i] > res {
			res = totalInformTime[i]
		}
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
