package double_pointers.valid_palindrome;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/valid-palindrome/">125. 验证回文串</a>
 * @since 2023-09-22 17:24
 */
class Solution {
    public boolean isPalindrome(String s) {
        s = s.toLowerCase();
        for (int i = 0, j = s.length() - 1; i < j;) {
            while (i < j && isNotValid(s.charAt(i))) {
                i++;
            }
            while (i < j && isNotValid(s.charAt(j))) {
                j--;
            }
            if (s.charAt(i) != s.charAt(j)) {
                return false;
            }
            i++;
            j--;
        }
        return true;
    }

    private boolean isNotValid(char c) {
        return (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') && (c < '0' || c > '9');
    }
}