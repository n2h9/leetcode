package solution

func findMaximumXOR(nums []int) int {
	trie := generateTrie(nums)
	offsprings := offs(trie)
	var num1, num2 *node
	currRes := 0
	if len(offsprings) <= 0 {
		return 0
	}
	if len(offsprings) == 1 {
		num1, num2 = offsprings[0], offsprings[0]
	} else {
		num1, num2 = offsprings[0], offsprings[1]
		currRes = 1
	}
	return maxXOR(num1, num2, currRes)
}

func maxXOR(num1, num2 *node, currRes int) int {
	offs1 := offs(num1)
	offs2 := offs(num2)
	if len(offs1) <= 0 || len(offs2) <= 0 {
		return currRes
	}
	if num1.offsprings[0] != nil && num2.offsprings[1] != nil ||
		num1.offsprings[1] != nil && num2.offsprings[0] != nil {
		val1 := 0
		if num1.offsprings[0] != nil && num2.offsprings[1] != nil {
			val1 = maxXOR(num1.offsprings[0], num2.offsprings[1], currRes<<1+1)
		}
		val2 := 0
		if num1.offsprings[1] != nil && num2.offsprings[0] != nil {
			val2 = maxXOR(num1.offsprings[1], num2.offsprings[0], currRes<<1+1)
		}
		return max(val1, val2)
	}

	val1 := 0
	if num1.offsprings[0] != nil && num2.offsprings[0] != nil {
		val1 = maxXOR(num1.offsprings[0], num2.offsprings[0], currRes<<1)
	}
	val2 := 0
	if num1.offsprings[1] != nil && num2.offsprings[1] != nil {
		val2 = maxXOR(num1.offsprings[1], num2.offsprings[1], currRes<<1)
	}
	return max(val1, val2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func offs(n *node) []*node {
	res := []*node{}
	for _, o := range n.offsprings {
		if o != nil {
			res = append(res, o)
		}
	}
	return res
}

type node struct {
	offsprings []*node
}

func newNode() *node {
	return &node{
		offsprings: make([]*node, 2),
	}
}

func generateTrie(nums []int) *node {
	trie := newNode()
	var i uint = 32
l:
	for {
		mask := 1 << i
		for j := 0; j < len(nums); j++ {
			if nums[j]&mask != 0 {
				break l
			}
		}
		if i == 0 {
			break
		}
		i--
	}

	for _, num := range nums {
		add(trie, num, i)
	}

	return trie
}

func add(trie *node, num int, highBitNum uint) *node {
	n := trie
	for {
		mask := 1 << highBitNum
		bit := num & mask >> highBitNum
		if n.offsprings[bit] == nil {
			n.offsprings[bit] = newNode()
		}
		n = n.offsprings[bit]
		if highBitNum == 0 {
			break
		}
		highBitNum--
	}

	return trie
}
