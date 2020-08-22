package solution

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	if len(flights) <= 0 {
		return -1
	}
	// [vertice] => [ [vertice, price] ]
	adjacentMap := make([][][]int, n)
	for _, flight := range flights {
		adjacentMap[flight[0]] = append(
			adjacentMap[flight[0]],
			[]int{flight[1], flight[2]},
		)
	}

	q := newPriorityQueue()
	q.push(
		newValue(src, 0, 0),
	)

	for !q.isEmpty() {
		node := q.top()
		q.pop()

		if node.vertice == dst {
			return node.cost
		}
		if node.stops > K {
			continue
		}

		for _, verticesWithPrice := range adjacentMap[node.vertice] {
			q.push(
				newValue(
					verticesWithPrice[0],
					node.stops+1,
					node.cost+verticesWithPrice[1],
				),
			)
		}
	}

	return -1
}
