package javasolutions;

import java.util.Arrays;

//https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements/
public class leetcode453 {
    public int minMoves(int[] nums) {
        Arrays.sort(nums);
        int sum = 0;
        for (int i = 1; i < nums.length; i++) {
            sum += nums[i] - nums[0];
        }
        return sum;
    }
}