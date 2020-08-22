package solution

func findDiagonalOrder(nums [][]int) []int {
	res := make([]int, 0, 100000)
	nextLevelAvailable, i, j := true, 0, 0
	for nextLevelAvailable {
		for ii, jj := i, j; ii >= 0; ii, jj = ii-1, jj+1 {
			if jj < len(nums[ii]) {
				res = append(res, nums[ii][jj])
			}
		}

		if (i + 1) < len(nums) {
			// j the same
			nextLevelAvailable, i = true, i+1
			continue
		}
		if (j + 1) < len(nums[i]) {
			// i the same
			nextLevelAvailable, j = true, j+1
			continue
		}

		nextLevelAvailable = false
		for ii, jj := i-1, j+1; ii >= 0; ii, jj = ii-1, jj+1 {
			if (jj + 1) < len(nums[ii]) {
				nextLevelAvailable, i, j = true, ii, jj+1
				break
			}
		}
	}

	return res
}
