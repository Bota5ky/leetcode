package simulation.image_smoother;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/image-smoother/">661. 图片平滑器</a>
 * @since 2023-09-07 14:10
 */
class Solution {
    public int[][] imageSmoother(int[][] M) {
        int[][] res = new int[M.length][M[0].length];
        for (int i = 0; i < M.length; i++) {
            for (int j = 0; j < M[i].length; j++) {
                res[i][j] = helper(M, i, j);
            }
        }
        return res;
    }

    private int helper(int[][] M, int i, int j) {
        int cnt = 0, sum = 0;
        int[][] position = {{i, j}, {i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}, {i - 1, j - 1},
                {i + 1, j - 1}, {i + 1, j + 1}, {i - 1, j + 1}};
        for (int[] ints : position) {
            if (ints[0] < 0 || ints[0] >= M.length || ints[1] < 0
                    || ints[1] >= M[0].length) {
                continue;
            }
            sum += M[ints[0]][ints[1]];
            cnt++;
        }
        return sum / cnt;
    }
}
