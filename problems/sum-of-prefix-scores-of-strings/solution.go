package solution

func sumPrefixScores(words []string) []int {
	trie := newNode()
	for _, w := range words {
		insert(trie, w)
	}
	res := make([]int, len(words))
	for i, w := range words {
		res[i] = sumValues(trie, w)
	}
	return res
}

type node struct {
	value    int
	children map[byte]*node
}

func newNode() *node {
	return &node{
		value:    0,
		children: make(map[byte]*node),
	}
}

func insert(n *node, s string) {
	for i := range s {
		if n.children[s[i]] == nil {
			n.children[s[i]] = newNode()
		}
		n = n.children[s[i]]
		n.value++
	}
}

func sumValues(n *node, s string) int {
	res := 0
	for i := range s {
		if n.children[s[i]] == nil {
			break // this situation should not be possible by the problem condition
		}
		n = n.children[s[i]]
		res += n.value
	}
	return res
}
