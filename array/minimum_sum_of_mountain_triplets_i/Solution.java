package array.minimum_sum_of_mountain_triplets_i;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/minimum-sum-of-mountain-triplets-i/">2908. 元素和最小的山形三元组 I</a>
 * @since 2024-03-29 21:38
 */
class Solution {
    public int minimumSum(int[] nums) {
        int sum = -1;
        int[] left = new int[nums.length];
        int[] right = new int[nums.length];
        left[0] = 51;
        right[nums.length - 1] = 51;
        for (int i = 1; i < nums.length; i++) {
            left[i] = Math.min(left[i - 1], nums[i - 1]);
        }
        for (int i = nums.length - 2; i >= 0; i--) {
            right[i] = Math.min(right[i + 1], nums[i + 1]);
        }
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] > left[i] && nums[i] > right[i]) {
                int temp = nums[i] + left[i] + right[i];
                if (sum < 0 || temp < sum) {
                    sum = temp;
                }
            }
        }
        return sum;
    }
}