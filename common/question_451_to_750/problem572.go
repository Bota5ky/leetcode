package question_451_to_750

// https://leetcode-cn.com/problems/subtree-of-another-tree/
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if t == nil {
		return true
	}
	if isSametree(s, t) {
		return true
	}
	if s != nil && (isSubtree(s.Left, t) || isSubtree(s.Right, t)) {
		return true
	}
	return false
}

func isSametree(s, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	if s.Val != t.Val {
		return false
	}
	return isSametree(s.Left, t.Left) && isSametree(s.Right, t.Right)
}
