package dynamic_programming.maximum_subarray;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/maximum-subarray/">53. 最大子数组和</a>
 * @since 2023-11-03 11:00
 */
class Solution {
    public int maxSubArray(int[] nums) {
        var sum = 0;
        var max = nums[0];
        for (int num : nums) {
            sum += num;
            max = Math.max(max, sum);
            if (sum < 0) {
                sum = 0;
            }
        }
        return max;
    }
}
