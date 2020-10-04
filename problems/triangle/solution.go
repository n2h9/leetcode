package solution

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n <= 0 {
		return 0
	}
	for i := n - 2; i >= 0; i-- {
		for j := len(triangle[i]) - 1; j >= 0; j-- {
			triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	return triangle[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
