/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/valid-sudoku/">36. 有效的数独</a>
 * @since 2024-01-29 11:24
 */
class Solution {
    public boolean isValidSudoku(char[][] board) {
        int[] row = new int[9];
        int[] column = new int[9];
        int[] block = new int[9];
        for (int i = 0; i < board.length; i++) {
            for (int j = 0; j < board[i].length; j++) {
                if (board[i][j] == '.') {
                    continue;
                }
                int p = i / 3 * 3 + j / 3;
                int val = 1 << (board[i][j] - '1');
                if ((row[i] & val) != 0 || (column[j] & val) != 0 || (block[p] & val) != 0) {
                    return false;
                }
                row[i] |= val;
                column[j] |= val;
                block[p] |= val;
            }
        }
        return true;
    }
}
