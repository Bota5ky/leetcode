package simulation.first_completely_painted_row_or_column;

import java.util.HashMap;
import java.util.Map;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/first-completely-painted-row-or-column/">2661. 找出叠涂元素</a>
 * @since 2024-02-06 20:10
 */
class Solution {
    public int firstCompleteIndex(int[] arr, int[][] mat) {
        Map<Integer, int[]> positions = new HashMap<>();
        int[] rows = new int[mat.length];
        int[] columns = new int[mat[0].length];
        for (int i = 0; i < mat.length; i++) {
            for (int j = 0; j < mat[i].length; j++) {
                positions.put(mat[i][j], new int[]{i, j});
            }
        }
        for (int i = 0; i < arr.length; i++) {
            int num = arr[i];
            int[] p = positions.get(num);
            int row = ++rows[p[0]];
            int column = ++columns[p[1]];
            if (row == mat[0].length || column == mat.length) {
                return i;
            }
        }
        return -1;
    }
}
