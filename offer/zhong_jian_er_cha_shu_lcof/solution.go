package zhong_jian_er_cha_shu_lcof

import . "leetcode/model"

// 剑指 Offer 07. 重建二叉树 https://leetcode.cn/problems/zhong-jian-er-cha-shu-lcof/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	mid := preorder[0]
	var head TreeNode
	head.Val = mid
	var i int
	for i = 0; inorder[i] != mid; i++ {
	}
	m := len(inorder[:i])
	n := len(inorder[i+1:])
	head.Left = buildTree(preorder[1:m+1], inorder[:i])
	head.Right = buildTree(preorder[m+1:m+n+1], inorder[i+1:])
	return &head
}
