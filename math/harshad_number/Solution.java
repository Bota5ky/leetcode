package math.harshad_number;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/harshad-number/">3099. 哈沙德数</a>
 * @since 2024-07-03 22:14
 */
class Solution {
    public int sumOfTheDigitsOfHarshadNumber(int x) {
        int digit = 0;
        int copy = x;
        while (copy > 0) {
            digit += copy % 10;
            copy /= 10;
        }
        return x % digit == 0 ? digit : -1;
    }
}
