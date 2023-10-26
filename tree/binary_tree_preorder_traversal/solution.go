package binary_tree_preorder_traversal

import . "leetcode/_model"

// 144. 二叉树的前序遍历 https://leetcode.cn/problems/binary-tree-preorder-traversal/
func preorderTraversal(root *TreeNode) []int {
	cur := root
	var res []int
	var stack []*TreeNode
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			res = append(res, cur.Val)
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1].Right
			stack = stack[:len(stack)-1]
		}
	}
	return res
}
