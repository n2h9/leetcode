package solution

func kClosest(points [][]int, K int) [][]int {
	// no need to calculate sqrt, because sqrt(x) < sqrt(y) <=> x < y
	dist := func(p []int) int {
		return p[0]*p[0] + p[1]*p[1]
	}
	// p1 is greater than p2 by priority if p1 is clother to 0,0 than p2
	gtByPrior := func(p1, p2 []int) bool {
		return dist(p1) < dist(p2)
	}
	swap := func(pp [][]int, i, j int) {
		pp[i], pp[j] = pp[j], pp[i]
	}
	var heapify func(pp [][]int, l, i int) 
	heapify = func(pp [][]int, l, i int) {
		if i >= l {
			return
		}
		mostPriority, left, right := i, 2*i+1, 2*i+2
		for _, child := range []int{left, right} {
			if child < l && gtByPrior(pp[child], pp[mostPriority]) {
				mostPriority = child
			}
		}
		if mostPriority != i {
			swap(pp, i, mostPriority)
			heapify(pp, l, mostPriority)
		}
	}
	var buble func(pp [][]int, i int)
	buble = func(pp [][]int, i int) {
		if i <= 0 {
			return
		}
		parent := (i-1)/2
		if gtByPrior(pp[i], pp[parent]) {
			swap(pp, i, parent)
			buble(pp, parent)
		}
	}
	buildHeap := func(pp [][]int) {
		l := len(pp)
		for i := 1; i < l; i++ {
			buble(pp, i)
		}
	}
	// -------------------------------------

	result := make([][]int, 0, K)

	buildHeap(points)
	heapLen := len(points)
	for ; K > 0; K-- {
		swap(points, 0, heapLen-1)
		result = append(result, points[heapLen-1])
		heapLen--
		heapify(points, heapLen, 0)
	}

    return result
}
