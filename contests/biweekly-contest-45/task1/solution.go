package solution

func sumOfUnique(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	sum := 0
	for k, v := range m {
		if v == 1 {
			sum += k
		}
	}
	return sum
}
