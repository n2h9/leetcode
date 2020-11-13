package solution

func subarraySum(nums []int, k int) int {
	n := len(nums)
	count := 0
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}
	return count
}
