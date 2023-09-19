package common.question_1_to_50.longest_common_prefix;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/longest-common-prefix/">14. 最长公共前缀</a>
 * @since 2023-09-19 22:48
 */
class Solution {
    public String longestCommonPrefix(String[] strs) {
        var sb = new StringBuilder();
        var sample = strs[0];
        for (int i = 0; i < strs[0].length(); i++) {
            var sampleChar = sample.charAt(i);
            for (int j = 1; j < strs.length; j++) {
                if (strs[j].length() <= i || strs[j].charAt(i) != sampleChar) {
                    return sb.toString();
                }
            }
            sb.append(sampleChar);
        }
        return sb.toString();
    }
}
