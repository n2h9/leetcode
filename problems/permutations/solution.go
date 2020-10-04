package solution

import "strconv"

func permute(nums []int) [][]int {
	if len(nums) <= 1 {
		return [][]int{
			nums,
		}
	}
	return backtracking(nums, make(map[string][][]int))
}

func backtracking(nums []int, store map[string][][]int) [][]int {
	hash := calcHash(nums)
	if res, ok := store[hash]; ok {
		return res
	}
	if len(nums) == 2 {
		res := [][]int{
			[]int{nums[0], nums[1]},
			[]int{nums[1], nums[0]},
		}
		store[hash] = res
		return res
	}
	res := make([][]int, 0)
	for i := range nums {
		b := make([]int, 0)
		b = append(b, nums[:i]...)
		b = append(b, nums[i+1:]...)
		pns := backtracking(b, store)
		for _, p := range pns {
			item := []int{nums[i]}
			item = append(item, p...)
			res = append(res, item)
		}
	}
	store[hash] = res
	return res
}

func calcHash(nums []int) string {
	hash := ""
	for _, v := range nums {
		hash += strconv.Itoa(v) + ","
	}
	return hash
}
