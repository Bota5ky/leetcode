package lcplcr.binary_tree_maximum_path_sum;


import model.TreeNode;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/jC7MId/">LCR 051. 二叉树中的最大路径和</a>
 * @since 2023-10-22 16:41
 */
class Solution {
    private int max;

    public int maxPathSum(TreeNode root) {
        max = root.val;
        maxPath(root);
        return max;
    }

    private int maxPath(TreeNode node) {
        if (node == null) {
            return 0;
        }
        var leftMax = Math.max(maxPath(node.left), 0);
        var rightMax = Math.max(maxPath(node.right), 0);
        max = Math.max(max, leftMax + rightMax + node.val);
        return node.val + Math.max(leftMax, rightMax);
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