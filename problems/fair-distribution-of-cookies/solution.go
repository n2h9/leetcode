package solution

func distributeCookies(cookies []int, k int) int {
	unfairness := 1<<32 - 1

	var determineUnfairness func(int, []int)
	determineUnfairness = func(index int, dist []int) {
		if index == len(cookies) {
			if localUnfairness := max(dist); localUnfairness <= unfairness {
				unfairness = localUnfairness
			}
			return
		}

		for i := 0; i < len(dist); i++ {
			dist[i] += cookies[index]
			determineUnfairness(index+1, dist)
			dist[i] -= cookies[index]
		}
	}

	determineUnfairness(0, make([]int, k))

	return unfairness
}

func max(nums []int) int {
	m := 0
	for _, v := range nums {
		if v > m {
			m = v
		}
	}
	return m
}
