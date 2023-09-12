package ping_heng_er_cha_shu_lcof

import . "leetcode/model"

// 剑指 Offer 55 - II. 平衡二叉树 https://leetcode.cn/problems/ping-heng-er-cha-shu-lcof/
// 面试题 04.04. 检查平衡性 https://leetcode.cn/problems/check-balance-lcci/
// 110. 平衡二叉树 https://leetcode.cn/problems/balanced-binary-tree/
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
