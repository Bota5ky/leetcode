package subtree_of_another_tree

import . "leetcode/_model"

// 572. 另一棵树的子树 https://leetcode.cn/problems/subtree-of-another-tree/
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if t == nil {
		return true
	}
	if isSameTree(s, t) {
		return true
	}
	if s != nil && (isSubtree(s.Left, t) || isSubtree(s.Right, t)) {
		return true
	}
	return false
}

func isSameTree(s, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	if s.Val != t.Val {
		return false
	}
	return isSameTree(s.Left, t.Left) && isSameTree(s.Right, t.Right)
}
