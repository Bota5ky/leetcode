package javasolutions;

//https://leetcode-cn.com/problems/image-smoother/
public class leetcode661 {
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
        int[][] position = { { i, j }, { i - 1, j }, { i + 1, j }, { i, j - 1 }, { i, j + 1 }, { i - 1, j - 1 },
                { i + 1, j - 1 }, { i + 1, j + 1 }, { i - 1, j + 1 } };
        for (int k = 0; k < position.length; k++) {
            if (position[k][0] < 0 || position[k][0] >= M.length || position[k][1] < 0
                    || position[k][1] >= M[0].length) {
                continue;
            }
            sum += M[position[k][0]][position[k][1]];
            cnt++;
        }
        return sum / cnt;
    }
}
