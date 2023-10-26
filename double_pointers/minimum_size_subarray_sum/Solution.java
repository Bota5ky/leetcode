package double_pointers.minimum_size_subarray_sum;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/minimum-size-subarray-sum/">209. 长度最小的子数组</a>
 * @since 2023-09-26 13:32
 */
class Solution {
    public int minSubArrayLen(int target, int[] nums) {
        int minLen = Integer.MAX_VALUE;
        var sum = nums[0];
        for (int i = 0, j = 0; ; ) {
            if (sum >= target) {
                minLen = Math.min(minLen, j - i + 1);
                if (minLen == 1) {
                    return 1;
                }
                sum -= nums[i++];
            } else {
                if (j == nums.length - 1) {
                    if (i == 0) {
                        return 0;
                    }
                    break;
                }
                sum += nums[++j];
            }
        }
        return minLen;
    }
}
