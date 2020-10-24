package solution

// func moveZeroes(nums []int) {
// 	for i := len(nums) - 1; i >= 0; i-- {
// 		if nums[i] == 0 {
// 			for j := i + 1; j < len(nums) && nums[j] != 0; j++ {
// 				nums[j-1], nums[j] = nums[j], nums[j-1]
// 			}
// 		}
// 	}
// }

// func moveZeroes(nums []int) {
// 	for i := len(nums) - 1; i >= 0; i-- {
// 		if nums[i] == 0 {
// 			nums = append(
// 				append(
// 					append(nums[:0], nums[:i]...),
// 					nums[i+1:]...,
// 				),
// 				0,
// 			)
// 		}
// 	}
// }

// func moveZeroes(nums []int) {
// 	zeroI, nonZeroI := 0, 0
// 	n := len(nums)
// 	for {
// 		// find first zero element index
// 		for ; zeroI < n && nums[zeroI] != 0; zeroI++ {
// 		}
// 		// stop if out of bound
// 		if zeroI >= n {
// 			break
// 		}
// 		nonZeroI = zeroI
// 		// find first non zero element
// 		for ; nonZeroI < n && nums[nonZeroI] == 0; nonZeroI++ {
// 		}
// 		// stop if out of bound
// 		if nonZeroI >= n {
// 			break
// 		}
// 		// swap
// 		nums[zeroI], nums[nonZeroI] = nums[nonZeroI], nums[zeroI]
// 	}
// }

// func moveZeroes(nums []int) {
// 	zeroI, nonZeroI := 0, 0
// 	n := len(nums)
// 	for {
// 		// find first zero element index
// 		for ; zeroI < n && nums[zeroI] != 0; zeroI++ {
// 		}
// 		// stop if out of bound
// 		if zeroI >= n {
// 			break
// 		}
// 		nonZeroI = zeroI
// 		// find first non zero element
// 		for ; nonZeroI < n && nums[nonZeroI] == 0; nonZeroI++ {
// 		}
// 		// stop if out of bound
// 		if nonZeroI >= n {
// 			break
// 		}
// 		// swap
// 		for zeroI < n && nonZeroI < n && nums[zeroI] == 0 && nums[nonZeroI] != 0 {
// 			nums[zeroI], nums[nonZeroI] = nums[nonZeroI], nums[zeroI]
// 			zeroI++
// 			nonZeroI++
// 		}
// 	}
// }

func moveZeroes(nums []int) {
	n := len(nums)
	i := 0
	for j := 0; j < n; j++ {
		if nums[j] != 0 {
			nums[i] = nums[j]
			i++
		}
	}
	for ; i < n; i++ {
		nums[i] = 0
	}
}
