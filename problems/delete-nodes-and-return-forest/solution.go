package solution

import (
	"sort"
)

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
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	return newProcessor(root, to_delete).run().result()
}

type processor struct {
	root      *TreeNode
	to_delete []int
	resultMap map[*TreeNode]bool
}

func newProcessor(root *TreeNode, to_delete []int) *processor {
	return &processor{
		root:      root,
		to_delete: to_delete,
		resultMap: make(map[*TreeNode]bool),
	}
}

func (p *processor) run() *processor {
	if p.root == nil {
		return p
	}
	p.resultMap[p.root] = true
	if len(p.to_delete) <= 0 {
		return p
	}
	sort.Ints(p.to_delete)
	p._delRec(p.root, nil)
	return p
}

func (p *processor) result() []*TreeNode {
	list := make([]*TreeNode, 0, len(p.resultMap))
	for k := range p.resultMap {
		list = append(list, k)
	}
	return list
}

func (p *processor) _delRec(node, parent *TreeNode) {
	if node == nil {
		return
	}
	next := []*TreeNode{node.Left, node.Right}

	if p._shouldDelete(node) {
		delete(p.resultMap, node)
		if parent != nil {
			if parent.Left == node {
				parent.Left = nil
			} else {
				parent.Right = nil
			}
		}
		if node.Left != nil {
			p.resultMap[node.Left] = true
			node.Left = nil
		}
		if node.Right != nil {
			p.resultMap[node.Right] = true
			node.Right = nil
		}
	}

	for _, n := range next {
		p._delRec(n, node)
	}
}

func (p *processor) _shouldDelete(node *TreeNode) bool {
	i := sort.SearchInts(p.to_delete, node.Val)
	return i < len(p.to_delete) && p.to_delete[i] == node.Val
}
