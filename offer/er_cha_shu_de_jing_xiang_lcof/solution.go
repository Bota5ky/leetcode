package er_cha_shu_de_jing_xiang_lcof

import . "leetcode/model"

// 剑指 Offer 27. 二叉树的镜像 https://leetcode.cn/problems/er-cha-shu-de-jing-xiang-lcof/
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
