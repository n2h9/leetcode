package solution

import (
	"crypto/md5"
	"strconv"
)

const checksumSize = md5.Size

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
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	return newProcessor().run(root).result()
}

type processor struct {
	resultList   []*TreeNode
	existenceMap map[[checksumSize]byte]int
}

func newProcessor() *processor {
	return &processor{
		resultList:   make([]*TreeNode, 0),
		existenceMap: make(map[[checksumSize]byte]int),
	}
}

func (p *processor) run(root *TreeNode) *processor {
	p.recHashDetermine(root)
	return p
}

func (p *processor) result() []*TreeNode {
	return p.resultList
}

// recursive calculate hash of left and right subtree
// calculate hash of node as hash (leftHash, value, rightHash)
// if hash is present in hashMap, add node to list
// otherwise add hash to hashMap
//
// return hash of node
func (p *processor) recHashDetermine(node *TreeNode) [checksumSize]byte {
	if node == nil {
		return [checksumSize]byte{}
	}
	hashLeft, hashRight := p.recHashDetermine(node.Left), p.recHashDetermine(node.Right)
	hashNode := hashAll(hashLeft, node.Val, hashRight)
	if p.existenceMap[hashNode] == 1 {
		p.resultList = append(p.resultList, node)
	}
	p.existenceMap[hashNode]++
	return hashNode
}

func hash(data []byte) [checksumSize]byte {
	return md5.Sum(data)
}

func hashAll(hash1 [checksumSize]byte, val int, hash2 [checksumSize]byte) [checksumSize]byte {
	data := append(
		append(hash1[:], []byte(strconv.Itoa(val))...),
		hash2[:]...,
	)
	return hash(data)
}
