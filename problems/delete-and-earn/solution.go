package solution

import (
	"sort"
)

func deleteAndEarn(nums []int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}
	sort.Ints(nums)
	distinct := make([]int, 2, n+1)
	sums := make([]int, 2, n+1)

	// start from 1 index because of dp formula depends dp(i-2)
	distinct[1], sums[1] = nums[0], nums[0]

	j := 1
	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1] {
			sums[j] += distinct[j]
			continue
		}
		distinct = append(distinct, nums[i])
		sums = append(sums, nums[i])
		j++
	}

	dp := make([]int, j+1)
	dp[1] = sums[1]
	for i := 2; i <= j; i++ {
		if distinct[i] > distinct[i-1]+1 {
			dp[i] = dp[i-1] + sums[i]
			continue
		}
		dp[i] = max(dp[i-1], dp[i-2]+sums[i])
	}

	return dp[j]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
