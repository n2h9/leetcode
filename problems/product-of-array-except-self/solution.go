package solution

func productExceptSelf(nums []int) []int {
	if len(nums) <= 1 {
		return []int{}
	}
	res := make([]int, len(nums))
	res[0] = 1
	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i-1]
	}
	acc := 1
	for i := len(nums) - 2; i >= 0; i-- {
		acc *= nums[i+1]
		res[i] *= acc
	}

	return res
}
