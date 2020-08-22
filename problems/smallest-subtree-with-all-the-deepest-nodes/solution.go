package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	levelMap := map[int]map[*TreeNode]bool{}
	level := createLevelMap(root, levelMap, 0)
	if len(levelMap[level]) <= 0 {
		return nil
	}
	if len(levelMap[level]) == 1 {
		for k := range levelMap[level] {
			return k
		}
	}
	node, _ := containsAllNodesNode(root, levelMap[level])
	return node
}

// return node which is an ancestor all nodes from a map
// or
// nill and number of nodes it is an ancestor
func containsAllNodesNode(
	root *TreeNode,
	nodeMap map[*TreeNode]bool,
) (*TreeNode, int) {
	if root == nil {
		return nil, 0
	}
	if nodeMap[root] {
		return nil, 1
	}
	nodeLeft, leftNum := containsAllNodesNode(root.Left, nodeMap)
	if nodeLeft != nil {
		return nodeLeft, leftNum
	}
	nodeRight, rightNum := containsAllNodesNode(root.Right, nodeMap)
	if nodeRight != nil {
		return nodeRight, rightNum
	}

	sum := leftNum + rightNum
	if sum >= len(nodeMap) {
		return root, sum
	}
	return nil, sum
}

// fill level map
// return max level
func createLevelMap(root *TreeNode, levelMap map[int]map[*TreeNode]bool, level int) int {
	if root == nil {
		return level - 1
	}
	if levelMap[level] == nil {
		levelMap[level] = make(map[*TreeNode]bool)
	}
	levelMap[level][root] = true
	ll := createLevelMap(root.Left, levelMap, level+1)
	lr := createLevelMap(root.Right, levelMap, level+1)
	if ll > lr {
		return ll
	}
	return lr
}
