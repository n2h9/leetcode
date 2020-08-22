package solution

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	n1, n2 := len(nums1), len(nums2)
	if n1 <= 0 || n2 <= 0 || k <= 0 {
		return [][]int{}
	}

	h := NewHeap(nil)
	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			h.Push(HeapItem{nums1[i], nums2[j]})
		}
	}

	res := make([][]int, 0, k)
	for ; k > 0 && !h.IsEmpty(); k-- {
		hi := h.Top()
		h.Pop()
		res = append(res, []int{hi.x, hi.y})
	}

	return res
}

type HeapItem struct {
	x int
	y int
}

func (hi *HeapItem) Priority() int {
	// -(x+y) because we need min heap
	return ^(hi.x + hi.y) + 1
}

type Heap struct {
	storage []HeapItem
}

func NewHeap(items []HeapItem) *Heap {
	return (&Heap{}).build(items)
}

func (h *Heap) Push(item HeapItem) {
	// this will be new last index
	i := len(h.storage)
	// add to the end and them move it upper if necessary
	h.storage = append(h.storage, item)
	p := (i - 1) / 2
	for i > 0 && h.storage[i].Priority() > h.storage[p].Priority() {
		h.storage[i], h.storage[p] = h.storage[p], h.storage[i]
		// and continue
		i = p
		p = (i - 1) / 2
	}
}

func (h *Heap) IsEmpty() bool {
	return len(h.storage) <= 0
}

func (h *Heap) Top() HeapItem {
	return h.storage[0]
}

// remove root element
func (h *Heap) Pop() {
	last := len(h.storage) - 1
	// move last to Top
	h.storage[0] = h.storage[last]
	// cut slice so that element is not accessable
	h.storage = h.storage[:last]
	// heapify root
	h.heapify(0)
}

func (h *Heap) build(items []HeapItem) *Heap {
	h.storage = make([]HeapItem, 0, len(items))
	for _, item := range items {
		h.Push(item)
	}
	return h
}

func (h *Heap) heapify(i int) {
	if i >= len(h.storage) {
		return
	}
	l := 2*i + 1
	r := 2*i + 2
	// index of max element
	mi := i
	if l < len(h.storage) && h.storage[l].Priority() > h.storage[mi].Priority() {
		mi = l
	}
	if r < len(h.storage) && h.storage[r].Priority() > h.storage[mi].Priority() {
		mi = r
	}
	if mi != i {
		// swap
		h.storage[i], h.storage[mi] = h.storage[mi], h.storage[i]
		h.heapify(mi)
	}
}
