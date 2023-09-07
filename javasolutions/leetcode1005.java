package javasolutions;

import java.util.Arrays;

public class leetcode1005 {
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
                nums[i] = - nums[i];
                k --;
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
