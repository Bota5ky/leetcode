package flatten_binary_tree_to_linked_list

import . "leetcode/_model"

// 114. 二叉树展开为链表 https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/
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
	root.Right = *pre
	*pre = root
}
