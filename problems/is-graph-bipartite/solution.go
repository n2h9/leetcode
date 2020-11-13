package solution

func isBipartite(graph [][]int) bool {
	const COLOR_YELLOW byte = 10
	const COLOR_BLUE byte = 20

	n := len(graph)

	colors := make([]byte, n)
	markWith := func(v int, color byte) {
		colors[v] = color
	}
	isMarked := func(v int) bool {
		return colors[v] != 0
	}

	var queue []int
	var qi int
	newQueue := func() (queue []int, qi int) {
		return make([]int, 0, n), 0
	}
	isEmpty := func() bool {
		return qi == len(queue)
	}
	top := func() int {
		return queue[qi]
	}
	pop := func() {
		qi++
	}
	push := func(v int) {
		queue = append(queue, v)
	}

	for i := 0; i < n; i++ {
		if isMarked(i) {
			continue
		}
		markWith(i, COLOR_YELLOW)
		queue, qi = newQueue()
		push(i)
		for !isEmpty() {
			v := top()
			pop()
			// switch color for marking adjacent vertexes
			color := COLOR_YELLOW
			if colors[v] == COLOR_YELLOW {
				color = COLOR_BLUE
			}
			for _, u := range graph[v] {
				if !isMarked(u) {
					markWith(u, color)
					push(u)
					continue
				}
				if colors[u] == colors[v] {
					// same color means that grapth is not bipartite
					return false
				}
			}
		}
	}

	return true
}
