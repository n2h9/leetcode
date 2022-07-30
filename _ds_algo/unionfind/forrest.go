package solution

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
