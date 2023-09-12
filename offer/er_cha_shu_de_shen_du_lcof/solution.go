package er_cha_shu_de_shen_du_lcof

import . "leetcode/model"

// 剑指 Offer 55 - I. 二叉树的深度 https://leetcode.cn/problems/er-cha-shu-de-shen-du-lcof/
// 104. 二叉树的最大深度 https://leetcode.cn/problems/maximum-depth-of-binary-tree/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return 1 + left
	}
	return 1 + right
}
