package solution

func kidsWithCandies(candies []int, extraCandies int) []bool {
	m := max(candies)
	r := make([]bool, len(candies))
	for i := 0; i < len(candies); i++ {
		r[i] = candies[i]+extraCandies >= m
	}

	return r
}

func max(a []int) int {
	m := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > m {
			m = a[i]
		}
	}

	return m
}
