package simulation.projection_area_of_3d_shapes;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/projection-area-of-3d-shapes/">883. 三维形体投影面积</a>
 * @since 2023-09-07 14:21
 */
class Solution {
    public int projectionArea(int[][] grid) {
        int[] rowMax = new int[grid.length];
        int[] colMax = new int[grid[0].length];
        int xy = 0;
        for (int i = 0; i < grid.length; i++) {
            for (int j = 0; j < grid[i].length; j++) {
                if (grid[i][j] > 0) {
                    xy++;
                }
                if (grid[i][j] > rowMax[i]) {
                    rowMax[i] = grid[i][j];
                }
                if (grid[i][j] > colMax[j]) {
                    colMax[j] = grid[i][j];
                }
            }
        }
        for (int max : rowMax) {
            xy += max;
        }
        for (int max : colMax) {
            xy += max;
        }
        return xy;
    }
}
