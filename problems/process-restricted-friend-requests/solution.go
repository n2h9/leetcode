package solution

func friendRequests(n int, restrictions [][]int, requests [][]int) []bool {
	forest := newForest(n)
	res := make([]bool, len(requests))

	for i, request := range requests {
		root0, root1 := forest.find(request[0]), forest.find(request[1])
		if root0 == root1 {
			res[i] = true
			continue
		}
		friendshipRestricted := false
		for _, restriction := range restrictions {
			root10, root11 := forest.find(restriction[0]), forest.find(restriction[1])
			if (root0 == root10 && root1 == root11) || (root0 == root11 && root1 == root10) {
				friendshipRestricted = true
				break
			}
		}

		if !friendshipRestricted {
			forest.union(request[0], request[1])
			res[i] = true
		}

	}

	return res
}

// https://en.wikipedia.org/wiki/Disjoint-set_data_structure
type ufforest struct {
	parent []int // [i] -> i-th parent
	rank   []int // [i] -> i-th rank
}

func newForest(size int) *ufforest {
	parent := make([]int, size)
	for i := 0; i < size; i++ {
		parent[i] = i
	}

	rank := make([]int, size)

	return &ufforest{
		parent: parent,
		rank:   rank,
	}
}

// returns x root
func (f *ufforest) find(x int) int {
	if f.parent[x] != x {
		f.parent[x] = f.find(f.parent[x])
		return f.parent[x]
	}
	return x
}

func (f *ufforest) union(x, y int) {
	x = f.find(x)
	y = f.find(y)

	if x == y {
		return
	}

	if f.rank[x] < f.rank[y] {
		x, y = y, x
	}

	f.parent[y] = x
	if f.rank[x] == f.rank[y] {
		f.rank[x]++
	}
}
