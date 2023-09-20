package common.question_1_to_50.trapping_rain_water;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/trapping-rain-water/">42. 接雨水</a>
 * @since 2023-09-19 15:29
 */
class Solution {
    public int trap(int[] height) {
        var maxLeft = new int[height.length];
        var maxRight = new int[height.length];
        for (int i = 1; i < height.length; i++) {
            maxLeft[i] = Math.max(maxLeft[i - 1], height[i - 1]);
        }
        for (int i = height.length - 2; i >= 0; i--) {
            maxRight[i] = Math.max(maxRight[i + 1], height[i + 1]);
        }
        var total = 0;
        for (int i = 1; i < height.length - 1; i++) {
            int edge = Math.min(maxLeft[i], maxRight[i]);
            if (height[i] < edge) {
                total += edge - height[i];
            }
        }
        return total;
    }
}
