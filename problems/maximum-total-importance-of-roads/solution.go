package solution

import "sort"

func maximumImportance(n int, roads [][]int) int64 {
	numOfConnectedPoints := make([]int, n)
	for _, road := range roads {
		numOfConnectedPoints[road[0]]++
		numOfConnectedPoints[road[1]]++
	}

	sort.Sort(sort.Reverse(sort.IntSlice(numOfConnectedPoints)))

	res := 0
	m := n
	for i := 0; i < n; i++ {
		res += numOfConnectedPoints[i] * m
		m--
	}

	return int64(res)
}
