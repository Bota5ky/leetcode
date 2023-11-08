package double_pointers.find_the_longest_balanced_substring_of_a_binary_string;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/find-the-longest-balanced-substring-of-a-binary-string/">2609. 最长平衡子字符串</a>
 * @since 2023-11-08 13:47
 */
class Solution {
    public int findTheLongestBalancedSubstring(String s) {
        int zeros = 0;
        int ones = 0;
        int len = 0;
        for (int i = 0; i < s.length(); i++) {
            if (ones > 0 && s.charAt(i) == '0') {
                len = Math.max(len, 2 * Math.min(zeros, ones));
                zeros = 1;
                ones = 0;
            } else {
                if (s.charAt(i) == '0') {
                    zeros++;
                } else {
                    ones++;
                }
            }
        }
        len = Math.max(len, 2 * Math.min(zeros, ones));
        return len;
    }
}
