package leetcode

//https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/
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
