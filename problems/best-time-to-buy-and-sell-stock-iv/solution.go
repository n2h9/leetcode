package solution

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n <= 0 {
		return 0
	}
	maxValue := 0
	dp := make([][]int, 2)
	dp[0] = make([]int, n)
	dp[1] = make([]int, n)

	for m := 0; m < k; m++ {
		currI := m % 2
		prevI := (m - 1) % 2
		dp[currI] = make([]int, n)
		for i := 2*m + 1; i < n; i++ {
			dp[currI][i] = orZero(dp, currI, i-1)
			for j := i - 1; j >= 2*m; j-- {
				dp[currI][i] = max(dp[currI][i], prices[i]-prices[j]+orZero(dp, prevI, j-1))
			}
		}
		maxValue = max(maxValue, dp[currI][n-1])
	}

	return maxValue
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func orZero(dp [][]int, m, i int) int {
	if i < 0 {
		return 0
	}
	if m < 0 {
		return 0
	}

	return dp[m][i]
}
