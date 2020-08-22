package solution

func finalPrices(prices []int) []int {
	result := make([]int, 0, len(prices))
	for i := range prices {
		res := prices[i]
		var j int
		for j = i + 1; j < len(prices) && prices[j] > prices[i]; j++ {
		}
		if j < len(prices) {
			res -= prices[j]
		}
		result = append(result, res)
	}
	return result
}
