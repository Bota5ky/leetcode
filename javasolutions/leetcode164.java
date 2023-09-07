package javasolutions;

import java.util.Arrays;

//https://leetcode-cn.com/problems/maximum-gap/
class leetcode164 {
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