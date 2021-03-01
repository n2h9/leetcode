package solution

func minPathSum(grid [][]int) int {
	min := func(x, y int) int {
		if y < x {
			return y
		}
		return x
	}

	m, n := len(grid), len(grid[0])

	top := make([]int, n)
	top[0] = grid[0][0]
	for i := 1; i < n; i++ {
		top[i] = top[i-1] + grid[0][i]
	}

	for i := 1; i < m; i++ {
		top[0] += grid[i][0]
		left := top[0]
		for j := 1; j < n; j++ {
			top[j] = min(top[j], left) + grid[i][j]
			left = top[j]
		}
	}

	return top[n-1]
}
