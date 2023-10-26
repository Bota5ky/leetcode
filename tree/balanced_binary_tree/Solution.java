package tree.balanced_binary_tree;

import _model.TreeNode;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/balanced-binary-tree/">110. 平衡二叉树</a>
 * @link <a href="https://leetcode.cn/problems/ping-heng-er-cha-shu-lcof/">剑指 Offer 55 - II. 平衡二叉树</a>
 * @link <a href="https://leetcode.cn/problems/check-balance-lcci/">面试题 04.04. 检查平衡性</a>
 * @link <a href="https://leetcode.cn/problems/ping-heng-er-cha-shu-lcof/">LCR 176. 判断是否为平衡二叉树</a>
 * @since 2023-10-19 22:01
 */
class Solution {
    public boolean isBalanced(TreeNode root) {
        return getHeight(root) >= 0;
    }

    private int getHeight(TreeNode node) {
        if (node == null) {
            return 0;
        }
        var left = getHeight(node.left);
        var right = getHeight(node.right);
        if (left < 0 || right < 0 || Math.abs(left - right) > 1) {
            return -1;
        }
        return Math.max(left, right) + 1;
    }
}

/*
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
