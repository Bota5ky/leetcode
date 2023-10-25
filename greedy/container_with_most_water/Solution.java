package greedy.container_with_most_water;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/container-with-most-water/">11. 盛最多水的容器</a>
 * @since 2023-09-22 22:24
 */
class Solution {
    public int maxArea(int[] height) {
        var max = 0;
        for (int i = 0, j = height.length - 1; i < j; ) {
            max = Math.max(max, Math.min(height[i], height[j]) * (j - i));
            if (height[i] <= height[j]) {
                i++;
            } else {
                j--;
            }
        }
        return max;
    }
}
