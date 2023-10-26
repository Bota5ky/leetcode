package string.reverse_words_in_a_string;

import java.util.Arrays;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/reverse-words-in-a-string/">151. 反转字符串中的单词</a>
 * @link <a href="https://leetcode.cn/problems/fan-zhuan-dan-ci-shun-xu-lcof/">LCR 181. 字符串中的单词反转</a>
 * @since 2023-09-20 09:33
 */
class Solution {
    public String reverseWords(String s) {
        var start = 0;
        var end = s.length();
        while (start < end && s.charAt(start) == ' ') {
            start++;
        }
        while (start < end && s.charAt(end - 1) == ' ') {
            end--;
        }

        var chars = new char[end - start];
        var j = 0;
        for (int i = end - 1; i >= start; i--) {
            char c = s.charAt(i);
            if (c == ' ' && i < s.length() - 1 && c == s.charAt(i + 1)) {
                continue;
            }
            chars[j++] = c;
        }
        for (int i = 0; i < j;) {
            var k = i + 1;
            while (k < j && chars[k] != ' ') {
                k++;
            }
            reverse(chars, i, k - 1);
            i = k + 1;
        }
        return String.valueOf(Arrays.copyOf(chars, j));
    }

    private void reverse(char[] chars, int start, int end) {
        for (int i = 0; i < (end - start + 1) / 2; i++) {
            var temp = chars[start + i];
            chars[start + i] =  chars[end - i];
            chars[end - i] = temp;
        }
    }
}
