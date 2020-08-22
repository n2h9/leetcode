package solution

import "sort"

func twoCitySchedCost(costs [][]int) int {
	sort.SliceStable(costs, func(i, j int) bool {
		difi := costs[i][0] - costs[i][1]
		if difi < 0 {
			difi = ^difi + 1
		}
		difj := costs[j][0] - costs[j][1]
		if difj < 0 {
			difj = ^difj + 1
		}
		if difi == difj {
			return costs[i][0] < costs[j][0]
		}
		return difi < difj
	})

	maxPerCity := len(costs) / 2
	sum := 0
	inA := 0
	inB := 0
	for i := len(costs) - 1; i >= 0; i-- {
		if inA >= maxPerCity {
			sum += costs[i][1]
			continue
		}
		if inB >= maxPerCity {
			sum += costs[i][0]
			continue
		}
		if costs[i][0] < costs[i][1] {
			sum += costs[i][0]
			inA++
			continue
		}
		if costs[i][0] > costs[i][1] {
			sum += costs[i][1]
			inB++
			continue
		}

		if inA < inB {
			sum += costs[i][0]
			inA++
		}

		sum += costs[i][1]
		inB++
	}
	return sum
}
