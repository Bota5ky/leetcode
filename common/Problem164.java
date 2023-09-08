package common;

import java.util.Arrays;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/maximum-gap/">164. 最大间距</a>
 * @since 2023-09-07 11:36
 */
class Problem164 {
    public int maximumGap(int[] nums) {
        Arrays.sort(nums);
        int gap = 0;
        for (int i = 1; i < nums.length; i++) {
            if (nums[i] - nums[i - 1] > gap) {
                gap = nums[i] - nums[i - 1];
            }
        }
        return gap;
    }
}
