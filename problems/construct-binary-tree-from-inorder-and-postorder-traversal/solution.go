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
func buildTree(inorder []int, postorder []int) *TreeNode {
	// assumed that len(inorder) always equal to len(postorder)
	if len(inorder) <= 0 {
		return nil
	}

	tree := &TreeNode{}
	tree.Val = postorder[len(postorder)-1]

	inOrderRootIndex := 0
	for ; inorder[inOrderRootIndex] != tree.Val; inOrderRootIndex++ {
	}

	tree.Left = buildTree(
		inorder[:inOrderRootIndex],
		postorder[:inOrderRootIndex],
	)

	tree.Right = buildTree(
		inorder[inOrderRootIndex+1:],
		postorder[inOrderRootIndex:len(postorder)-1],
	)

	return tree
}
