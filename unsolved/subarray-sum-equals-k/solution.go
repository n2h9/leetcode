package solution

func subarraySum(nums []int, k int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}
	dp := make([]map[int]int, n)
	kmin, kmax := -1000*n, n*1000
	dp[0] = make(map[int]int)
	dp[0][nums[0]] = 1
	for i := 1; i < n; i++ {
		dp[i] = make(map[int]int)
		dp[i][nums[i]] = 1
		for j := kmin; j <= kmax; j++ {
			dp[i][j] += dp[i-1][j-nums[i]]
		}
	}

	count := 0
	for i := 0; i < n; i++ {
		count += dp[i][k]
	}
	return count
}
