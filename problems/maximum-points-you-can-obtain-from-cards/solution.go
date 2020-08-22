package solution

func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)

	sum := 0
	for i := n - k; i < n; i++ {
		sum += cardPoints[i]
	}
	max := sum

	for i := 0; i < k; i++ {
		sum -= cardPoints[n-k+i]
		sum += cardPoints[i]
		if sum > max {
			max = sum
		}
	}

	return max
}
