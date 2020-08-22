package solution

func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	maxProfit := 0
	minPrice := prices[0]
	for i := 1; i < n; i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
			continue
		}
		if p := prices[i] - minPrice; p > maxProfit {
			maxProfit = p
		}
	}

	return maxProfit
}
