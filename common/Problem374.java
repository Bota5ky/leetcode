package common;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/guess-number-higher-or-lower/">374. 猜数字大小</a>
 * @since 2023-09-07 13:53
 */
class Problem374 {
    public int guessNumber(int n) {
        for (int i = 1, j = n; i <= j; ) {
            int mid = i + (j - i) / 2;
            if (guess(mid, n) < 0) {
                j = mid - 1;
            } else if (guess(mid, n) > 0) {
                i = mid + 1;
            } else {
                return mid;
            }
        }
        return -1;
    }

    int guess(int num, int n) {
        return Integer.compare(num, n);
    }
}
