package solution

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	minCost := make([]int, n)
	minCost[0] = cost[0]
	minCost[1] = cost[1]

	for i := 2; i < n; i++ {
		minCost[i] = min(minCost[i-1], minCost[i-2]) + cost[i]
	}

	return min(minCost[n-1], minCost[n-2])
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
