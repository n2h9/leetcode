package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	vals := money(root)
	return max(vals[0], vals[1])
}

// return
// [0] => max rob value if node is robbed
// [1] => max rob value if node is not robbed
func money(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}

	lm := money(root.Left)
	rm := money(root.Right)

	return []int{
		root.Val + lm[1] + rm[1],
		max(lm[0], lm[1]) + max(rm[0], rm[1]),
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
