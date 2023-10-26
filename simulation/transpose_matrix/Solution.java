package simulation.transpose_matrix;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/transpose-matrix/">867. 转置矩阵</a>
 * @since 2023-09-07 14:19
 */
class Solution {
    public int[][] transpose(int[][] matrix) {
        int m = matrix.length, n = matrix[0].length;
        int[][] transposed = new int[n][m];
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                transposed[j][i] = matrix[i][j];
            }
        }
        return transposed;
    }
}
