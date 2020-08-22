package solution

func twoSum(nums []int, target int) []int {
	n := len(nums)
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[j]+nums[i] == target {
				return []int{j, i}
			}
		}
	}
	return nil
}
