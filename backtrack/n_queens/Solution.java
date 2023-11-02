package backtrack.n_queens;

import java.util.*;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/n-queens/">51. N 皇后</a>
 * @since 2023-11-02 21:30
 */
class Solution {
    public List<List<String>> solveNQueens(int n) {
        var res = new ArrayList<List<String>>();
        var queens = new int[n];
        Arrays.fill(queens, -1);
        backtrack(res, queens, 0, new HashSet<>(), new HashSet<>(), new HashSet<>());
        return res;
    }

    private void backtrack(List<List<String>> res, int[] queens, int row, Set<Integer> columns, Set<Integer> diagonals1, Set<Integer> diagonals2) {
        if (row == queens.length) {
            res.add(generateBoard(queens));
            return;
        }
        for (int i = 0; i < queens.length; i++) {
            if (columns.contains(i)) {
                continue;
            }
            if (diagonals1.contains(row - i)) {
                continue;
            }
            if (diagonals2.contains(row + i)) {
                continue;
            }
            queens[row] = i;
            columns.add(i);
            diagonals1.add(row - i);
            diagonals2.add(row + i);
            backtrack(res, queens, row + 1, columns, diagonals1, diagonals2);
            queens[row] = -1;
            columns.remove(i);
            diagonals1.remove(row - i);
            diagonals2.remove(row + i);
        }
    }

    private List<String> generateBoard(int[] queens) {
        var board = new ArrayList<String>(queens.length);
        for (int queen : queens) {
            var row = new char[queens.length];
            Arrays.fill(row, '.');
            row[queen] = 'Q';
            board.add(new String(row));
        }
        return board;
    }
}
