package solution

func findMaxForm(strs []string, m int, n int) int {
	l := len(strs)
	if l <= 0 || m < 0 || n < 0 {
		return 0
	}
	dp := make([][][]int, l+1)
	// init zero border cases for convenient usage
	dp[0] = make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[0][i] = make([]int, n+1)
	}
	// --
	for k := 1; k <= l; k++ {
		dp[k] = make([][]int, m+1)
		z, o := numOfZerosOnes(strs[k-1])
		for i := 0; i <= m; i++ {
			dp[k][i] = make([]int, n+1)
			for j := 0; j <= n; j++ {
				dp[k][i][j] = dp[k-1][i][j]
				if z <= i && o <= j {
					dp[k][i][j] = max(dp[k][i][j], dp[k-1][i-z][j-o]+1)
				}
			}
		}
	}
	return dp[l][m][n]
}

func numOfZerosOnes(a string) (int, int) {
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] == '0' {
			count++
		}
	}
	return count, len(a) - count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
