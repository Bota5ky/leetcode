package maximize_sum_of_array_after_k_negations

import java.util.Arrays;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/maximize-sum-of-array-after-k-negations/">1005. K 次取反后最大化的数组和</a>
 * @since 2023-09-07 14:23
 */
class Problem1005 {
    public int largestSumAfterKNegations(int[] nums, int k) {
        Arrays.sort(nums);
        int minAbs = Math.abs(nums[0]);
        for (int i = 0; i < nums.length && k > 0; i++) {
            if (Math.abs(nums[i]) <= minAbs) {
                minAbs = Math.abs(nums[i]);
            } else {
                break;
            }
            if (nums[i] < 0) {
                nums[i] = -nums[i];
                k--;
            }
        }
        if (k % 2 == 1) {
            nums[0] -= 2 * minAbs;
        }
        int sum = 0;
        for (int num : nums) {
            sum += num;
        }
        return sum;
    }
}
