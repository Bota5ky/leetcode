package math.find_the_divisibility_array_of_a_string;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/find-the-divisibility-array-of-a-string/">2575. 找出字符串的可整除数组</a>
 * @since 2024-03-07 08:34
 */
class Solution {
    public int[] divisibilityArray(String word, int m) {
        int[] div = new int[word.length()];
        long sum = 0;
        for (int i = 0; i < word.length(); i++) {
            int num = word.charAt(i) - '0';
            sum = sum * 10 + num;
            sum %= m;
            if (sum == 0) {
                div[i] = 1;
            }
        }
        return div;
    }
}
