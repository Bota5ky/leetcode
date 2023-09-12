package first_common_ancestor_lcci

import . "leetcode/model"

// 面试题 04.08. 首个共同祖先 https://leetcode.cn/problems/first-common-ancestor-lcci/
// 236. 二叉树的最近公共祖先 https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/
func lowestCommonAncestor3(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p {
		return p
	}
	if root == q {
		return q
	}
	left := lowestCommonAncestor3(root.Left, p, q)
	right := lowestCommonAncestor3(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}
