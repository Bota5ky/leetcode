package two_pointers.find_the_index_of_the_first_occurrence_in_a_string;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/">28. 找出字符串中第一个匹配项的下标</a>
 * @since 2023-09-21 22:43
 */
class Solution {
    // KMP算法
    public int strStr(String haystack, String needle) {
        var next = buildNext(needle);
        for (int i = 0, j = 0; i < haystack.length(); i++) {
            while (j > 0 && haystack.charAt(i) != needle.charAt(j)) {
                j = next[j - 1];
            }
            if (haystack.charAt(i) == needle.charAt(j)) {
                j++;
            }
            if (j == needle.length()) {
                return i - j + 1;
            }
        }
        return -1;
    }

    private int[] buildNext(String pattern) {
        var next = new int[pattern.length()];
        for (int i = 1, j = 0; i < pattern.length(); i++) {
            while (j > 0 && pattern.charAt(i) != pattern.charAt(j)) {
                j = next[j - 1];
            }
            if (pattern.charAt(i) == pattern.charAt(j)) {
                j++;
            }
            next[i] = j;
        }
        return next;
    }
}
