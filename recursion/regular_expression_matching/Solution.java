package recursion.regular_expression_matching;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/regular-expression-matching/">10. 正则表达式匹配</a>
 * @since 2023-11-03 11:26
 */
class Solution {
    public boolean isMatch(String s, String p) {
        if (p.isEmpty()) {
            return s.isEmpty();
        }
        // first match
        if (!s.isEmpty() && (s.charAt(0) == p.charAt(0) || p.charAt(0) == '.')) {
            if (p.length() > 1 && p.charAt(1) == '*') {
                return isMatch(s.substring(1), p) || isMatch(s, p.substring(2));
            }
            return isMatch(s.substring(1), p.substring(1));
        }
        // first not match
        return p.length() >= 2 && p.charAt(1) == '*' && isMatch(s, p.substring(2));
    }
}
