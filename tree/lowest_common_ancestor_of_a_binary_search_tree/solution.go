package lowest_common_ancestor_of_a_binary_search_tree

import . "leetcode/_model"

// 235. 二叉搜索树的最近公共祖先 https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree/
// LCR 193. 二叉搜索树的最近公共祖先 https://leetcode.cn/problems/er-cha-sou-suo-shu-de-zui-jin-gong-gong-zu-xian-lcof/
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
