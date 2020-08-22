package solution

func rotate(nums []int, k int) {
	n := len(nums)
	if n < 2 {
		return
	}

	// when k == n it means that array stays the same
	k = k % n
	if k <= 0 {
		return
	}

	// O(1) space
	_rotate2(nums, k)
}

// O(n) time; o(k%n) space
func _rotate1(nums []int, k int) {
	n := len(nums)
	carriers := make([]int, 0, k)
	for i := n - k; i < n; i++ {
		carriers = append(carriers, nums[i])
	}

	for i, j := 0, 0; i < n; i++ {
		nums[i], carriers[j] = carriers[j], nums[i]
		j++
		if j >= k {
			j = 0
		}
	}
}

// O((k%n)*n) time; o(1) space
func _rotate2(nums []int, k int) {
	n := len(nums)
	for j := 0; j < k; j++ {
		carrier := nums[n-1]
		for i := 0; i < n; i++ {
			nums[i], carrier = carrier, nums[i]
		}
	}
}
