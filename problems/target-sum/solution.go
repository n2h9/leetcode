package solution

func findTargetSumWays(nums []int, S int) int {
	n := len(nums)
	if n <= 0 || S > 1000 {
		return 0
	}

	dp := make([][]int, n+1)
	for i := 0; i < n; i++ {
		// by problem condition sum of elements in nums <= 1000
		// taken 1000 for negative sum and 1000 for positive sum and 1 for zero sum
		dp[i] = make([]int, 2001)
	}
	dp[n] = make([]int, 2001)
	// how many ways exist to create sum using first number
	initNum := 1
	if nums[0] == 0 {
		initNum = 2
	}
	dp[1][convertIndex(-nums[0])] = initNum
	dp[1][convertIndex(nums[0])] = initNum

	for k := 2; k <= n; k++ {
		for s := -1000; s <= 1000; s++ {
			prev1 := 0
			if s-nums[k-1] >= -1000 {
				prev1 = dp[k-1][convertIndex(s-nums[k-1])]
			}
			prev2 := 0
			if s+nums[k-1] <= 1000 {
				prev2 = dp[k-1][convertIndex(s+nums[k-1])]
			}
			dp[k][convertIndex(s)] = prev1 + prev2
		}
	}

	return dp[n][convertIndex(S)]
}

func convertIndex(s int) int {
	return s + 1000
}
