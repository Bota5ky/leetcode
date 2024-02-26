package tree.range_sum_of_bst;

import _model.TreeNode;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/range-sum-of-bst/">938. 二叉搜索树的范围和</a>
 * @since 2024-02-26 08:13
 */
class Solution {
    public int rangeSumBST(TreeNode root, int low, int high) {
        if (root == null) {
            return 0;
        }
        if (root.val < low) {
            return rangeSumBST(root.right, low, high);
        }else if (root.val > high) {
            return rangeSumBST(root.left, low, high);
        }
        return root.val + rangeSumBST(root.left, low, high) + rangeSumBST(root.right, low, high);
    }
}
