package solution

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countPairs(root *TreeNode, distance int) int {
	_, total := countPairsRec(root, distance)
	return total
}

// return
// leafs map [len of path from leave to root] => number of leaves
func countPairsRec(root *TreeNode, dist int) (leafs map[int]int, total int) {
	if root == nil {
		return nil, 0
	}
	leftLeafs, leftTotal := countPairsRec(root.Left, dist)
	rightLeafs, rightTotal := countPairsRec(root.Right, dist)
	if leftLeafs == nil && rightLeafs == nil {
		// leaf root
		return map[int]int{1: 1}, 0
	}

	// determine number of good leafs
	// it is possible if and only if both subtre have leafs
	total = leftTotal + rightTotal
	if leftLeafs != nil && rightLeafs != nil {
		for leftLen, leftNum := range leftLeafs {
			for rightLen, rightNum := range rightLeafs {
				if leftLen+rightLen <= dist {
					total += leftNum * rightNum
				}
			}
		}
	}

	// merge leafs maps
	mergedLeafs := leftLeafs
	if mergedLeafs == nil {
		mergedLeafs = rightLeafs
	} else {
		for k, v := range rightLeafs {
			mergedLeafs[k] += v
		}
	}
	// --

	// increase path len by 1 level
	leafs = make(map[int]int, len(mergedLeafs))
	for k, v := range mergedLeafs {
		leafs[k+1] = v
	}

	return leafs, total
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
