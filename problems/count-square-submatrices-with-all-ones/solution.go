package solution

func countSquares(matrix [][]int) int {
	n := len(matrix)
	if n <= 0 {
		return 0
	}
	m := len(matrix[0])
	if m <= 0 {
		return 0
	}
	// dp - number of square matrix for subnatrix ends in i,j
	// (indexes from 1 to n and from 1 to m)
	// 00
	// ends is a number of square submatrix of all ones EXACTLY ENDS in i,j
	// (or same as len of max square submatrix  of all one ends in i,j)
	dp, ends := make([][]int, n+1), make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i], ends[i] = make([]int, m+1), make([]int, m+1)
	}
	dp[1][1] = matrix[0][0]

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if v := matrix[i-1][j-1]; v == 1 {
				ends[i][j] = 1 + min(ends[i-1][j-1], ends[i-1][j], ends[i][j-1])
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1] + ends[i][j]
		}
	}

	return dp[n][m]
}

func min(x ...int) int {
	m := x[0]
	for i := 1; i < len(x); i++ {
		if x[i] < m {
			m = x[i]
		}
	}
	return m
}
