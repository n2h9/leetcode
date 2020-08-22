package solution

import "encoding/json"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	s := nodeToSerializable(root)
	b, e := json.Marshal(s)
	if e != nil {
		return ""
	}
	return string(b)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) <= 0 {
		return nil
	}
	s := &TreeNodeSerializable{}
	if e := json.Unmarshal([]byte(data), s); e != nil {
		return nil
	}

	return serializableToTreeNode(s)
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeNodeSerializable struct {
	Val   int                   `json:"v"`
	Left  *TreeNodeSerializable `json:"l"`
	Right *TreeNodeSerializable `json:"r"`
}

func nodeToSerializable(node *TreeNode) *TreeNodeSerializable {
	if node == nil {
		return nil
	}

	return &TreeNodeSerializable{
		Val:   node.Val,
		Left:  nodeToSerializable(node.Left),
		Right: nodeToSerializable(node.Right),
	}
}

func serializableToTreeNode(node *TreeNodeSerializable) *TreeNode {
	if node == nil {
		return nil
	}

	return &TreeNode{
		Val:   node.Val,
		Left:  serializableToTreeNode(node.Left),
		Right: serializableToTreeNode(node.Right),
	}
}
