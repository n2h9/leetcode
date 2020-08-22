package solution

func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = 0
	for i := 1; i < n; i++ {
		p := 0
		for j := i - 1; j >= 0; j-- {
			dpi := j - 2
			if dpi < 0 {
				dpi = 0
			}
			p = max(p, prices[i]-prices[j]+dp[dpi])
		}
		dp[i] = max(dp[i-1], p)
	}

	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
