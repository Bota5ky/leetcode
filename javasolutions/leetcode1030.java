package javasolutions;

import java.util.Arrays;
import java.util.Comparator;

//https://leetcode-cn.com/problems/matrix-cells-in-distance-order/
class leetcode1030 {
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
        Arrays.sort(res, new Comparator<int[]>() {
            public int compare(int[] o1, int[] o2) {
                return (Math.abs(o1[0] - r0) + Math.abs(o1[1] - c0)) - (Math.abs(o2[0] - r0) + Math.abs(o2[1] - c0));
            }
        });
        // System.out.println(res);
        return res;
    }
    
    public static void main(String[] args) {
        leetcode1030 sol = new leetcode1030();
        sol.allCellsDistOrder(1, 2, 0, 0);
    }
}