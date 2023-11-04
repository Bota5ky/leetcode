package simulation.queens_that_can_attack_the_king;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/queens-that-can-attack-the-king/">1222. 可以攻击国王的皇后</a>
 * @since 2023-09-14 10:29
 */
class Solution {
    public List<List<Integer>> queensAttacktheKing(int[][] queens, int[] king) {
        ArrayList<List<Integer>> res = new ArrayList<List<Integer>>();
        boolean[][] exist = new boolean[8][8];
        for (int[] queen : queens) {
            exist[queen[0]][queen[1]] = true;
        }
        // left
        for (int i = king[0]; i >= 0; i--) {
            if (exist[i][king[1]]) {
                res.add(Arrays.asList(i, king[1]));
                break;
            }
        }
        // right
        for (int i = king[0]; i < 8; i++) {
            if (exist[i][king[1]]) {
                res.add(Arrays.asList(i, king[1]));
                break;
            }
        }
        // up
        for (int i = king[1]; i < 8; i++) {
            if (exist[king[0]][i]) {
                res.add(Arrays.asList(king[0], i));
                break;
            }
        }
        // down
        for (int i = king[1]; i >= 0; i--) {
            if (exist[king[0]][i]) {
                res.add(Arrays.asList(king[0], i));
                break;
            }
        }
        // left up
        for (int i = 1; king[0] - i >= 0 && king[1] + i < 8; i++) {
            if (exist[king[0] - i][king[1] + i]) {
                res.add(Arrays.asList(king[0] - i, king[1] + i));
                break;
            }
        }
        // right up
        for (int i = 1; king[0] + i < 8 && king[1] + i < 8; i++) {
            if (exist[king[0] + i][king[1] + i]) {
                res.add(Arrays.asList(king[0] + i, king[1] + i));
                break;
            }
        }
        // left down
        for (int i = 1; king[0] - i >= 0 && king[1] - i >= 0; i++) {
            if (exist[king[0] - i][king[1] - i]) {
                res.add(Arrays.asList(king[0] - i, king[1] - i));
                break;
            }
        }
        // right down
        for (int i = 1; king[0] + i < 8 && king[1] - i >= 0; i++) {
            if (exist[king[0] + i][king[1] - i]) {
                res.add(Arrays.asList(king[0] + i, king[1] - i));
                break;
            }
        }
        return res;
    }
}
