package solution

func maxProfit(prices []int) int {
	n := len(prices)
	if n <= 0 {
		return 0
	}

	// max profit for 1 transaction
	dp1 := make([]int, n)
	// max profit for 2 transaction
	dp2 := make([]int, n)

	// first day we could only buy or do nothing so no profit at first day
	dp1[0], dp2[0] = 0, 0

	// tranasction 1 process
	for i := 1; i < n; i++ {
		m := 0
		for j := 0; j < i; j++ {
			if p := prices[i] - prices[j]; p > m {
				m = p
			}
		}
		dp1[i] = max(dp1[i-1], m)
	}

	// transaction 2 process
	// day 0 - buy
	// day 1 - sell
	// day 2 - buy
	// day 3 - sell - so 3 is min day index we could sell a stock on second transcation
	for i := 3; i < n; i++ {
		m := 0
		// day 0 - buy
		// day 1 - sell
		// day 2 - buy again => so 2 is min day index for second transaction buy
		for j := 2; j < i; j++ {
			if p := prices[i] - prices[j] + dp1[j-1]; p > m {
				m = p
			}
		}
		dp2[i] = max(dp2[i-1], m)
	}

	return max(dp1[n-1], dp2[n-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
