package tree.house_robber_iii;

import _model.TreeNode;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/house-robber-iii/">337. 打家劫舍 III</a>
 * @since 2023-09-18 09:51
 */
class Solution {
    public int rob(TreeNode root) {
        int[] maxRob = dfs(root);
        return Math.max(maxRob[0], maxRob[1]);
    }

    private int[] dfs(TreeNode root) {
        if (root == null) {
            return new int[]{0, 0};
        }
        int[] r = dfs(root.right);
        int[] l = dfs(root.left);
        int robCurrent = root.val + r[1] + l[1];
        int notRobCurrent = Math.max(l[0], l[1]) + Math.max(r[0], r[1]);
        return new int[]{robCurrent, notRobCurrent};
    }
}
