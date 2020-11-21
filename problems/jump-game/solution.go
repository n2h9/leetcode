package solution

func canJump(nums []int) bool {
	n := len(nums)
	if n <= 0 {
		return false
	}
	// dp[i] determines if i-th element is accessable from 0 element
	dp := make([]bool, n)
	dp[0] = true
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if !dp[j] {
				continue
			}
			if j + nums[j] >= i {
				dp[i] = true
				break
			}
		}
	}
	return dp[n-1]   
}
