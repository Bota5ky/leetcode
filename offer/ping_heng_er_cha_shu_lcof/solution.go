package ping_heng_er_cha_shu_lcof

import . "leetcode/model"

// 剑指 Offer 55 - II. 平衡二叉树 https://leetcode.cn/problems/ping-heng-er-cha-shu-lcof/
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := depthOfTreeNode(root.Left)
	right := depthOfTreeNode(root.Right)
	if left-right > 1 || left-right < -1 {
		return false
	}
	return isBalanced(root.Right) && isBalanced(root.Left)
}

func depthOfTreeNode(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := depthOfTreeNode(root.Left)
	right := depthOfTreeNode(root.Right)
	if left >= right {
		return 1 + left
	}
	return 1 + right
}
