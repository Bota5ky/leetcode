package math.maximum_swap;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/maximum-swap/">670. 最大交换</a>
 * @since 2024-02-04 21:14
 */
class Solution {
    public int maximumSwap(int num) {
        if (num < 10) {
            return num;
        }
        int[] bit = new int[9];
        int i = 0;
        while (num != 0) {
            bit[i++] = num % 10;
            num /= 10;
        }
        int[] max = new int[9];
        for (i = 1; i < bit.length; i++) {
            if (bit[i] > bit[max[i - 1]]) {
                max[i] = i;
            } else {
                max[i] = max[i - 1];
            }
        }
        i = 8;
        while (bit[i] == 0) {
            i--;
        }
        for (; i > 0; i--) {
            if (bit[i] < bit[max[i - 1]]) {
                int temp = bit[max[i - 1]];
                bit[max[i - 1]] = bit[i];
                bit[i] = temp;
                break;
            }
        }
        int sum = 0;
        i = 8;
        while (bit[i] == 0) {
            i--;
        }
        for (; i >= 0; i--){
            sum = sum * 10 + bit[i];
        }
        return sum;
    }
}
