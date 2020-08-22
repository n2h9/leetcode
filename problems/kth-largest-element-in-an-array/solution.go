package solution

// k is always valid by condition 1 <= k <= len(nums)
func findKthLargest(nums []int, k int) int {
	n := len(nums)
	if n <= 0 {
		return -1
	}
	// after buildHeap max element will be at nums[0]
	// so for k = 1 answer is found
	buildHeap(nums, 1)
	for i := 1; i < k; i++ {
		// move max element to the end of array, and do not take it to account anymore
		nums[0], nums[n-i] = nums[n-i], nums[0]
		// push max element to the start of array
		heapify(nums, 0, n-1-i)
	}
	return nums[0]
}

// recurrently creates a heap from slice
// call buildHeap(a, 0)
func buildHeap(a []int, ind int) {
	if ind >= len(a) {
		return
	}

	// for only one element no actions
	if ind == 0 {
		buildHeap(a, 1)
		return
	}

	i := ind
	p := (ind - 1) / 2
	// while a[i] gt parent move it up
	for i > 0 && a[i] > a[p] {
		a[i], a[p] = a[p], a[i]
		i = p
		p = (i - 1) / 2
	}

	buildHeap(a, ind+1)
}

func heapify(a []int, ind, last int) {
	// max index, left index, right index
	mi, li, ri := ind, 2*ind+1, 2*ind+2
	// determine element with max index
	for _, i := range []int{li, ri} {
		if i <= last && a[i] > a[mi] {
			mi = i
		}
	}
	if mi == ind {
		// top element is max nothing to do
		return
	}
	a[mi], a[ind] = a[ind], a[mi]
	heapify(a, mi, last)
}
