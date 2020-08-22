package solution

type HeapItem interface {
	Priority() int
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
	// move last to top
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
