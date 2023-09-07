package javasolutions;

//https://leetcode-cn.com/problems/guess-number-higher-or-lower/
public class leetcode374 {
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
