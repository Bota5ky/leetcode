package question_201_to_300

// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p {
		return p
	}
	if root == q {
		return q
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}
