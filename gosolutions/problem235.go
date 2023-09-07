package gosolutions

// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//保证p在左 q在右
	if p.Val > q.Val {
		p, q = q, p
	}
	//在根节点的左右子树
	if p.Val <= root.Val && q.Val >= root.Val {
		return root
	}
	//都在左边
	if p.Val <= root.Val && q.Val <= root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	return lowestCommonAncestor(root.Right, p, q)
}
