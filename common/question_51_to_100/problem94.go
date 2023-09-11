package question_51_to_100

import "leetcode/common/question_1_to_50"

// https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
func inorderTraversal(root *common.TreeNode) []int {
	cur := root
	var res []int
	var stack []*common.TreeNode
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, node.Val)
			cur = node.Right
		}
	}
	return res
}