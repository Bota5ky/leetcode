package tree.lowest_common_ancestor_of_a_binary_search_tree;

import _model.TreeNode;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/er-cha-sou-suo-shu-de-zui-jin-gong-gong-zu-xian-lcof/">LCR 193. 二叉搜索树的最近公共祖先</a>
 * @since 2023-10-25 17:01
 */
class Solution {
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        if (root == null || (root.val <= p.val && root.val >= q.val) || (root.val >= p.val && root.val <= q.val)) {
            return root;
        }
        if (root.val > p.val) {
            return lowestCommonAncestor(root.left, p ,q);
        }
        return lowestCommonAncestor(root.right, p, q);
    }
}
