package common;

import java.util.Arrays;
import java.util.Comparator;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/matrix-cells-in-distance-order/">1030. 距离顺序排列矩阵单元格</a>
 * @since 2023-09-07 14:25
 */
class Problem1030 {
    public int[][] allCellsDistOrder(int R, int C, int r0, int c0) {
        int[][] res = new int[R * C][2];
        int k = 0;
        for (int i = 0; i < R; i++) {
            for (int j = 0; j < C; j++) {
                res[k][0] = i;
                res[k][1] = j;
                k++;
            }
        }
        Arrays.sort(res, Comparator.comparingInt(o -> (Math.abs(o[0] - r0) + Math.abs(o[1] - c0))));
        return res;
    }
}
