package common.question_751_to_1000;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/projection-area-of-3d-shapes/">883. 三维形体投影面积</a>
 * @since 2023-09-07 14:21
 */
class Problem883 {
    public int projectionArea(int[][] grid) {
        int[] rowmax = new int[grid.length];
        int[] colmax = new int[grid[0].length];
        int xy = 0;
        for (int i = 0; i < grid.length; i++) {
            for (int j = 0; j < grid[i].length; j++) {
                if (grid[i][j] > 0) {
                    xy++;
                }
                if (grid[i][j] > rowmax[i]) {
                    rowmax[i] = grid[i][j];
                }
                if (grid[i][j] > colmax[j]) {
                    colmax[j] = grid[i][j];
                }
            }
        }
        for (int i = 0; i < rowmax.length; i++) {
            xy += rowmax[i];
        }
        for (int i = 0; i < colmax.length; i++) {
            xy += colmax[i];
        }
        return xy;
    }
}
