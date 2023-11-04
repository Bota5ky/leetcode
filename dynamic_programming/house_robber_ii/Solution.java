package common.question_201_to_300.house_robber_ii;

import java.util.Arrays;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/house-robber-ii/">213. 打家劫舍 II</a>
 * @since 2023-09-17 08:39
 */
class Solution {
    public int rob(int[] nums) {
        if (nums.length == 1) {
            return nums[0];
        }
        return Math.max(stole(Arrays.copyOf(nums, nums.length - 1)),
                stole(Arrays.copyOfRange(nums, 1, nums.length)));
    }

    private int stole(int[] nums) {
        int[] dp = new int[nums.length + 2];
        for (int i = 0; i < nums.length; i++) {
            dp[i + 2] = Math.max(dp[i + 1], dp[i] + nums[i]);
        }
        return dp[dp.length - 1];
    }
}