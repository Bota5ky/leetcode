package same_tree

import "leetcode/common/question_1_to_50"

// https://leetcode.cn/problems/same-tree/
func isSameTree(p *common.TreeNode, q *common.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val == q.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return false
}
