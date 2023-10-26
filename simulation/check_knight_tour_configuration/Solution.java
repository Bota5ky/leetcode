package simulation.check_knight_tour_configuration;

import java.util.HashMap;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/check-knight-tour-configuration/">2596. 检查骑士巡视方案</a>
 * @since 2023-09-13 15:42
 */
class Solution {
    public boolean checkValidGrid(int[][] grid) {
        if (grid[0][0] != 0) {
            return false;
        }
        int steps = grid.length * grid.length;
        var move = new HashMap<Integer, int[]>(steps);
        for (int i = 0; i < grid.length; i++) {
            for (int j = 0; j < grid[i].length; j++) {
                move.put(grid[i][j], new int[]{i, j});
            }
        }
        for (int i = 1; i < steps; i++) {
            var pre = move.get(i - 1);
            var cur = move.get(i);
            if (Math.abs((pre[0] - cur[0]) * (pre[1] - cur[1])) != 2) {
                return false;
            }
        }
        return true;
    }
}
