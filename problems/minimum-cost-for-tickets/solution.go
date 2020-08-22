package solution

func mincostTickets(days []int, costs []int) int {
	dp := make([]int, 366)
	day := 1

	for i := 0; i < len(days); i++ {
		// pass through days we do not plan voyage
		for ; day < days[i]; day++ {
			dp[day] = dp[day-1]
		}
		dp[day] = min(
			valOrZero(dp, day-1)+costs[0],
			valOrZero(dp, day-7)+costs[1],
			valOrZero(dp, day-30)+costs[2],
		)
		day++
	}

	return dp[day-1]
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

func valOrZero(a []int, i int) int {
	if i < 0 {
		return 0
	}

	return a[i]
}
