package tree.inorder_successor_in_bst;

import _model.TreeNode;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/inorder-successor-in-bst/">285. 二叉搜索树中的中序后继</a>
 * @link <a href="https://leetcode.cn/problems/P5rCT8/">LCR 053. 二叉搜索树中的中序后继</a>
 * @since 2023-11-08 15:03
 */
class Solution {
    public TreeNode inorderSuccessor(TreeNode root, TreeNode p) {
        if (root == null) {
            return null;
        }
        if (root.val > p.val) {
            TreeNode left = inorderSuccessor(root.left, p);
            if (left != null) {
                return left;
            }
            return root;
        }
        TreeNode right = inorderSuccessor(root.right, p);
        if (right == null && root.val > p.val) {
            return root;
        }
        return right;
    }
}
