package solution

import (
	"sort"
)

type KthLargest struct {
	k          int
	sortedList *node
}

func Constructor(k int, nums []int) KthLargest {
	sort.Ints(nums)
	return KthLargest{
		k:          k,
		sortedList: newFromSortedList(nums),
	}
}

func (this *KthLargest) Add(val int) int {
	this.sortedList = addNode(this.sortedList, newNode(val))
	node := this.sortedList
	for i := this.k - 1; i > 0; i-- {
		node = node.next
	}
	return node.value
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

type node struct {
	value int
	next  *node
}

func newNode(value int) *node {
	return &node{value: value}
}

func newFromSortedList(nums []int) *node {
	i := len(nums) - 1
	if i < 0 {
		return nil
	}
	list := newNode(nums[i])
	curr := list
	for i := i - 1; i >= 0; i-- {
		curr.next = newNode(nums[i])
		curr = curr.next
	}

	return list
}

func addNode(list *node, node *node) *node {
	if list == nil {
		return node
	}

	if node.value >= list.value {
		node.next = list
		return node
	}

	curr := list.next
	prev := list
	for curr != nil && curr.value > node.value {
		prev = curr
		curr = curr.next
	}

	prev.next = node
	node.next = curr

	return list
}
