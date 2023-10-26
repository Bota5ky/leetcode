package maximum_depth_of_binary_tree

import . "leetcode/_model"

// 104. 二叉树的最大深度 https://leetcode.cn/problems/maximum-depth-of-binary-tree/
// LCR 175. 计算二叉树的深度 https://leetcode.cn/problems/er-cha-shu-de-shen-du-lcof/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return maxInt(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
