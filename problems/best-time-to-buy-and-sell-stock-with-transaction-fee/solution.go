package solution

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	prevBuy, buy := 0, 0-prices[0]-fee
	prevSell, sell := 0, 0
	for i := 1; i < n; i++ {
		prevBuy, prevSell = buy, sell
		buy = max(prevBuy, prevSell-prices[i]-fee)
		sell = max(prevSell, prevBuy+prices[i])
	}
	return sell
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
