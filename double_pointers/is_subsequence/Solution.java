package double_pointers.is_subsequence;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/is-subsequence/">392. 判断子序列</a>
 * @since 2023-09-22 17:40
 */
class Solution {
    public boolean isSubsequence(String s, String t) {
        if (s.length() > t.length()) {
            return false;
        }
        int i = 0;
        for (int j = 0; i < s.length() && j < t.length();) {
            if (s.charAt(i) == t.charAt(j)) {
                i++;
            }
            j++;
        }
        return i == s.length();
    }
}
