package shu_de_zi_jie_gou_lcof

import . "leetcode/model"

// 剑指 Offer 26. 树的子结构 https://leetcode.cn/problems/shu-de-zi-jie-gou-lcof/
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	//遍历判断是否是相同节点
	if A == nil {
		return B == nil
	}
	if B == nil {
		return false
	}
	//子树是否相同
	if isSameTree(A, B) {
		return true
	}
	return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func isSameTree(A *TreeNode, B *TreeNode) bool {
	if A == nil {
		return B == nil
	}
	if B == nil {
		return true
	}
	if A.Val != B.Val {
		return false
	}
	return isSameTree(A.Left, B.Left) && isSameTree(A.Right, B.Right)
}
