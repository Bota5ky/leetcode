package dynamic_programming.house_robber;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/house-robber/">198. 打家劫舍</a>
 * @since 2023-09-16 09:57
 */
class Solution {
    public int rob(int[] nums) {
        int[] dp = new int[nums.length + 2];
        for (int i = 0; i < nums.length; i++) {
            dp[i + 2] = Math.max(dp[i] + nums[i], dp[i + 1]);
        }
        return dp[dp.length - 1];
    }
}
