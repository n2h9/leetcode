package solution

func jump(nums []int) int {
	const MAX_INT = 1<<31
	n := len(nums)
	// dp[i] number of min jump to reach i from 0
	dp := make([]int, n)
	dp[0] = 0
	for i := 1; i < n; i++ {		
		val := MAX_INT
		for j := 0; j < i; j++ {			
			if j + nums[j] >= i && dp[j] < val {
				val = dp[j]
			}
		}
		dp[i] = val+1
	}
	return dp[n-1]   
}
