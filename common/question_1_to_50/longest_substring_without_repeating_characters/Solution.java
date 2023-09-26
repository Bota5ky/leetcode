package common.question_1_to_50.longest_substring_without_repeating_characters;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/longest-substring-without-repeating-characters/"></a>
 * @since 2023-09-26 14:09
 */
class Solution {
    public int lengthOfLongestSubstring(String s) {
        var record = new int[128];
        var maxLen = 0;
        for (int i = 0, j = 0; j < s.length(); j++) {
            char idx = s.charAt(j);
            if (record[idx] > 0) {
                var p = record[idx];
                for (int k = i; k < record[idx]; k++) {
                    record[s.charAt(k)] = 0;
                }
                i = p;
            }
            record[idx] = j + 1;
            maxLen = Math.max(maxLen, j - i + 1);
        }
        return maxLen;
    }
}
