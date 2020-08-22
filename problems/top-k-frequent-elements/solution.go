package solution

// Given a non-empty array of integers,
// You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
// --
// do not check nums and k validity
func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int)
	for _, num := range nums {
		countMap[num]++
	}

	hp := newHeap()
	for value, priority := range countMap {
		hp.push(
			&heapItem{
				value:    value,
				priority: priority,
			},
		)
	}

	result := make([]int, 0, k)
	for ; k > 0 && !hp.isEmpty(); k-- {
		result = append(result, hp.top().value)
		hp.pop()
	}
	return result
}

type heapItem struct {
	value    int
	priority int
}

type heap struct {
	storage []*heapItem
}

func newHeap() *heap {
	return &heap{
		storage: make([]*heapItem, 0),
	}
}

func (h *heap) isEmpty() bool {
	return len(h.storage) <= 0
}

func (h *heap) push(item *heapItem) {
	h.storage = append(h.storage, item)
	ind := len(h.storage) - 1
	// parent index
	pi := (ind - 1) / 2
	// move item up until it is > then parent
	for ind > 0 && h.storage[ind].priority > h.storage[pi].priority {
		h.storage[ind], h.storage[pi] = h.storage[pi], h.storage[ind]
		ind = pi
		pi = (ind - 1) / 2
	}
}

func (h *heap) top() *heapItem {
	return h.storage[0]
}

func (h *heap) pop() {
	h.storage[0], h.storage[len(h.storage)-1] = h.storage[len(h.storage)-1], h.storage[0]
	h.storage = h.storage[:len(h.storage)-1]
	h._heapify(0)
}

func (h *heap) _heapify(ind int) {
	// left index, right index, max index
	li, ri, mi := 2*ind+1, 2*ind+2, ind

	for _, i := range []int{li, ri} {
		if i < len(h.storage) && h.storage[i].priority > h.storage[mi].priority {
			mi = i
		}
	}

	if mi != ind {
		h.storage[mi], h.storage[ind] = h.storage[ind], h.storage[mi]
		h._heapify(mi)
	}
}
