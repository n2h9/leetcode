package solution

func maxAbsoluteSum(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	abs := func(x int) int {
		if x >= 0 {
			return x
		}
		return ^x + 1
	}

	max := nums[0]
	sum, sum2 := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		if nums[i] > sum {
			sum = nums[i]
		}
		sum2 += nums[i]
		if nums[i] < sum2 {
			sum2 = nums[i]
		}
		if abs(sum) > abs(max) {
			max = sum
		}
		if abs(sum2) > abs(max) {
			max = sum2
		}
	}

	return abs(max)
}
