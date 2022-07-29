package solution

func countPairs(n int, edges [][]int) int64 {
	adjList := make([][]int, n)
	connected := make([]int, n)
	for i := 0; i < n; i++ {
		adjList[i] = make([]int, 0)
		connected[i] = -1
	}

	for _, edge := range edges {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}

	for i := 0; i < n; i++ {
		if connected[i] != -1 {
			continue // node already belongs to some cluster
		}
		stack := newStack()
		stack.push(i)
		for !stack.isEmpty() {
			node := stack.pop()
			connected[node] = i
			for _, child := range adjList[node] {
				if connected[child] != -1 {
					continue // already processed
				}
				stack.push(child)
			}
		}
	}

	sizes := make(map[int]int)
	for _, cluster := range connected {
		sizes[cluster]++
	}

	sum := 0
	nodesNum := 0

	for _, size := range sizes {
		sum += size * nodesNum
		nodesNum += size
	}

	return int64(sum)
}

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
