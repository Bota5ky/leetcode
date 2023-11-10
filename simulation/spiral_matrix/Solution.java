package simulation.spiral_matrix;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/">LCR 146. 螺旋遍历二维数组</a>
 * @since 2023-11-10 14:16
 */
class Solution {
    public int[] spiralArray(int[][] array) {
        int up = 0;
        int down = array.length - 1;
        int left = 0;
        int columnSize = array.length == 0 ? 0 : array[0].length;
        int right = columnSize - 1;
        int d = 0;
        int[] iterator = new int[array.length * columnSize];
        int i = 0;
        while (down >= up && left <= right) {
            if (d == 0) {
                for (int j = left; j <= right; j++) {
                    iterator[i++] = array[up][j];
                }
                up++;
            } else if (d == 1) {
                for (int j = up; j <= down; j++) {
                    iterator[i++] = array[j][right];
                }
                right--;
            } else if (d == 2) {
                for (int j = right; j >= left; j--) {
                    iterator[i++] = array[down][j];
                }
                down--;
            } else {
                for (int j = down; j >= up; j--) {
                    iterator[i++] = array[j][left];
                }
                left++;
            }
            d = (d + 1) % 4;
        }
        return iterator;
    }
}
