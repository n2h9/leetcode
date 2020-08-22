package solution

func maxSubArray(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	max := nums[0]
	sum := nums[0]

	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		if nums[i] > sum {
			sum = nums[i]
		}
		if sum > max {
			max = sum
		}
	}

	return max
}
