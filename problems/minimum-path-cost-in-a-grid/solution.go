package solution

func minPathCost(grid [][]int, moveCost [][]int) int {
	// max signed int value for 32 bit system
	maxInt := 1<<31 - 1

	// init dp array
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = maxInt
		}
	}

	// set min cost to reach the topest row equal to the top row values
	for i := 0; i < n; i++ {
		dp[0][i] = grid[0][i]
	}

	// do main run to determine min cost for each cell sequentially
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if cost := dp[i-1][k] + moveCost[grid[i-1][k]][j] + grid[i][j]; cost < dp[i][j] {
					dp[i][j] = cost
				}
			}
		}
	}

	// find the min value from the bottom line of dp array
	min := maxInt
	for i := 0; i < n; i++ {
		if dp[m-1][i] < min {
			min = dp[m-1][i]
		}
	}

	return min
}
