package common

// 654321后序遍历
// https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/
func flatten(root *TreeNode) {
	var pre *TreeNode
	connect(root, &pre)
}
func connect(root *TreeNode, pre **TreeNode) {
	if root == nil {
		return
	}
	connect(root.Right, pre)
	connect(root.Left, pre)
	root.Left = nil
	root.Right = (*pre)
	*pre = root
}

// TreeNode 二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
