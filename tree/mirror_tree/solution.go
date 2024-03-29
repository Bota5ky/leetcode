package er_cha_shu_de_jing_xiang_lcof

import . "leetcode/_model"

// LCR 144. 翻转二叉树 https://leetcode.cn/problems/er-cha-shu-de-jing-xiang-lcof/
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	temp := root.Left
	root.Left = root.Right
	root.Right = temp
	mirrorTree(root.Left)
	mirrorTree(root.Right)
	return root
}
