package tree.lowest_common_ancestor_of_a_binary_tree;

import model.TreeNode;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/er-cha-shu-de-zui-jin-gong-gong-zu-xian-lcof/">LCR 194. 二叉树的最近公共祖先</a>
 * @since 2023-10-25 16:42
 */
class Solution {
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        if (root == null || root == p || root == q) {
            return root;
        }
        var left = lowestCommonAncestor(root.left, p, q);
        var right = lowestCommonAncestor(root.right, p, q);
        if (left != null && right != null) {
            return root;
        }
        return left == null ? right : left;
    }
}
