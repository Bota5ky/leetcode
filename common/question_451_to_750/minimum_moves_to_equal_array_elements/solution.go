package common.question_451_to_750;

import java.util.Arrays;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/minimum-moves-to-equal-array-elements/">453. 最小操作次数使数组元素相等</a>
 * @since 2023-09-07 14:05
 */
class Problem453 {
    public int minMoves(int[] nums) {
        Arrays.sort(nums);
        int sum = 0;
        for (int i = 1; i < nums.length; i++) {
            sum += nums[i] - nums[0];
        }
        return sum;
    }
}
