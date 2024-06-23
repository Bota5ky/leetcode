package string.detect_capital;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/detect-capital/">520. 检测大写字母</a>
 * @since 2024-06-23 09:15
 */
class Solution {
    public boolean detectCapitalUse(String word) {
        char firstChar = word.charAt(0);
        if (isLower(firstChar)) {
            for (int i = 1; i < word.length(); i++) {
                char c = word.charAt(i);
                if (!isLower(c)) {
                    return false;
                }
            }
        } else {
            if (word.length() == 1) {
                return true;
            }
            boolean second = isLower(word.charAt(1));
            for (int i = 2; i < word.length(); i++) {
                if (second != isLower(word.charAt(i))) {
                    return false;
                }
            }
        }
        return true;
    }

    private boolean isLower(char c) {
        return c >= 'a' && c <= 'z';
    }
}
