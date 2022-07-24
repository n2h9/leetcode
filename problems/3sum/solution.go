package solution

import "sort"

func threeSum(nums []int) [][]int {
	n := len(nums)
	triplets := newContainer()
	distinctValues := make(map[int]map[int]bool, n)

	for i := 0; i < n; i++ {
		if distinctValues[nums[i]] == nil {
			distinctValues[nums[i]] = make(map[int]bool, n)
		}
		distinctValues[nums[i]][i] = true
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			sum := nums[i] + nums[j]
			if distinctValues[(-1)*sum] != nil {
				hasAnotherIndex := false
				for k := range distinctValues[(-1)*sum] {
					if k != i && k != j {
						hasAnotherIndex = true
						break
					}
				}
				if !hasAnotherIndex {
					continue
				}
				triplet := []int{(-1) * sum, nums[i], nums[j]}
				sort.Ints(triplet)
				if triplets.exist(triplet) {
					continue
				}
				triplets.put(triplet)
			}
		}
	}

	return triplets.get()
}

type container struct {
	triplets  [][]int
	existence map[int]map[int]map[int]bool
}

func newContainer() *container {
	return &container{
		triplets:  make([][]int, 0, 32),
		existence: make(map[int]map[int]map[int]bool),
	}
}

func (c *container) get() [][]int {
	return c.triplets
}

func (c *container) exist(triplet []int) bool {
	if c.existence[triplet[0]] == nil {
		return false
	}
	if c.existence[triplet[0]][triplet[1]] == nil {
		return false
	}
	return c.existence[triplet[0]][triplet[1]][triplet[2]]
}

func (c *container) put(triplet []int) {
	if c.existence[triplet[0]] == nil {
		c.existence[triplet[0]] = make(map[int]map[int]bool)
	}
	if c.existence[triplet[0]][triplet[1]] == nil {
		c.existence[triplet[0]][triplet[1]] = make(map[int]bool)
	}

	c.existence[triplet[0]][triplet[1]][triplet[2]] = true

	c.triplets = append(c.triplets, triplet)
}
