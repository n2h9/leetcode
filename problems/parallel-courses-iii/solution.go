package solution

func minimumTime(n int, relations [][]int, time []int) int {
	childrenList := make([][]int, n+1)
	parentsList := make([][]int, n+1)

	for _, relation := range relations {
		childrenList[relation[0]] = append(childrenList[relation[0]], relation[1])
		parentsList[relation[1]] = append(parentsList[relation[1]], relation[0])
	}

	q := newQueue()

	// put in a queue courses which not depend on any course
	for i := 1; i <= n; i++ {
		if len(parentsList[i]) <= 0 {
			q.push(i)
		}
	}

	totalTime := make([]int, n+1)
	processed := make([]bool, n+1)

	for !q.isEmpty() {
		node := q.pop()
		if processed[node] {
			continue
		}
		notAllParantsProcessed := false
		for _, parent := range parentsList[node] {
			if !processed[parent] {
				notAllParantsProcessed = true
				break
			}
		}
		if notAllParantsProcessed {
			continue
		}

		// add own time to total time
		totalTime[node] += time[node-1]

		for _, child := range childrenList[node] {
			if totalTime[node] > totalTime[child] {
				totalTime[child] = totalTime[node]
			}
			q.push(child)
		}
		processed[node] = true
	}

	max := totalTime[1]
	for i := 2; i <= n; i++ {
		if totalTime[i] > max {
			max = totalTime[i]
		}
	}

	return max
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
