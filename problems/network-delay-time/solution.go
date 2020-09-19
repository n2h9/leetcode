package solution

func networkDelayTime(times [][]int, N int, K int) int {
	al := make([][][]int, N+1)
	for _, edge := range times {
		al[edge[0]] = append(al[edge[0]], []int{edge[1], edge[2]})
	}
	if !isAllConnected(al, N, K) {
		return -1
	}

	// Bellman-Ford
	cost := make([]int, N+1)
	for i := range cost {
		cost[i] = -1
	}
	cost[K] = 0
	for i := 1; i < N; i++ {
		for _, edge := range times {
			src, dst, weight := edge[0], edge[1], edge[2]
			if cost[src] != -1 {
				if newCost := cost[src] + weight; cost[dst] == -1 || newCost < cost[dst] {
					cost[dst] = newCost
				}
			}
		}
	}
	// --

	max := 0
	for _, v := range cost {
		if v > max {
			max = v
		}
	}

	return max
}

func isAllConnected(al [][][]int, n, k int) bool {
	visited := make([]bool, n+1)
	return countVertex(al, k, visited) == n
}

func countVertex(al [][][]int, vertex int, visited []bool) int {
	visited[vertex] = true
	count := 1
	for _, edge := range al[vertex] {
		if v := edge[0]; !visited[v] {
			count += countVertex(al, v, visited)
		}
	}
	return count
}
