package solution

func minHeightShelves(books [][]int, shelf_width int) int {
	n := len(books)
	if n == 1 {
		return books[0][1]
	}

	dp := make([]int, n+1)
	dp[0] = 0

	for i := 1; i <= n; i++ {
		h := books[i-1][1]
		w := books[i-1][0]
		dp[i] = dp[i-1] + h
		for j := i - 1; j > 0 && w+books[j-1][0] <= shelf_width; j-- {
			h = max(h, books[j-1][1])
			w += books[j-1][0]
			dp[i] = min(dp[i], h+dp[j-1])
		}
	}

	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
