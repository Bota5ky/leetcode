package common;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/super-pow/">372. 超级次方</a>
 * @since 2023-09-07 13:48
 */
class Problem372 {
    public int superPow(int a, int[] b) {
        int result = 1;
        for (int i = b.length - 1; i >= 0; i--) {
            result = result * smartPow(a, b[i]) % 1337;
            a =smartPow(a, 10);
        }
        return result;
    }

    public int smartPow(int a, int b) {
        a = a % 1337;
        int result = 1;
        for (int i = 0; i < b; i++) {
            result = result * a % 1337;
        }
        return result;
    }
}
