package array.max_increase_to_keep_city_skyline;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/max-increase-to-keep-city-skyline/">807. 保持城市天际线</a>
 * @since 2024-07-14 20:53
 */
class Solution {
    public int maxIncreaseKeepingSkyline(int[][] grid) {
        int[] rowMax = new int[grid.length];
        int[] colMax = new int[grid[0].length];
        for (int i = 0; i < grid.length; i++) {
            for (int j = 0; j < grid[i].length; j++) {
                rowMax[i] = Math.max(rowMax[i], grid[i][j]);
                colMax[j] = Math.max(colMax[j], grid[i][j]);
            }
        }
        int sum = 0;
        for (int i = 0; i < grid.length; i++) {
            for (int j = 0; j < grid[i].length; j++) {
                sum += Math.min(rowMax[i], colMax[j]) - grid[i][j];
            }
        }
        return sum;
    }
}
