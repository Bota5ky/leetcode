package double_pointers.max_consecutive_ones_iii;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/max-consecutive-ones-iii/">1004. 最大连续1的个数 III</a>
 * @since 2023-11-03 21:49
 */
class Solution {
    public int longestOnes(int[] nums, int k) {
        int maxOnes = 0;
        int zeros = 0;
        for (int i = 0, j = 0; j < nums.length; j++) {
            if (nums[j] == 0) {
                zeros++;
            }
            if (zeros > k) {
                while (nums[i] != 0) {
                    i++;
                }
                zeros--;
                i++;
            }
            maxOnes = Math.max(maxOnes, j - i + 1);
        }
        return maxOnes;
    }
}
