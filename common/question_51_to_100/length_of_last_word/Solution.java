package common.question_51_to_100.length_of_last_word;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/length-of-last-word/">58. 最后一个单词的长度</a>
 * @since 2023-09-19 22:42
 */
class Solution {
    public int lengthOfLastWord(String s) {
        var i = s.length() - 1;
        while (i >= 0 && s.charAt(i) == ' ') {
            i--;
        }
        var j = i - 1;
        while (j >= 0 && s.charAt(j) != ' ') {
            j--;
        }
        return i - j;
    }
}