package string.string_to_integer_atoi;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/string-to-integer-atoi/">8. 字符串转换整数 (atoi)</a>
 * @since 2023-11-09 13:59
 */
class Solution {
    public int myAtoi(String s) {
        int num = 0;
        int flag = 1;
        int i = 0;
        while (i < s.length() && s.charAt(i) == ' ') {
            i++;
        }
        if (i < s.length() && (s.charAt(i) == '-' || s.charAt(i) == '+')) {
            flag = s.charAt(i) == '-' ? -1 : 1;
            i++;
        }
        while (i < s.length() && s.charAt(i) == '0') {
            i++;
        }
        int edge = i + 9;
        for (; i < s.length() && s.charAt(i) >= '0' && s.charAt(i) <= '9'; i++) {
            int currentNumber = s.charAt(i) - '0';
            if (i == edge) {
                int number = num * flag;
                if (number > Integer.MAX_VALUE / 10 || number == Integer.MAX_VALUE / 10 && currentNumber >= Integer.MAX_VALUE % 10) {
                    return Integer.MAX_VALUE;
                }
                if (number < Integer.MIN_VALUE / 10 || number == Integer.MIN_VALUE / 10 && flag * currentNumber <= Integer.MIN_VALUE % 10) {
                    return Integer.MIN_VALUE;
                }
            } else if (i > edge){
                return flag == 1 ? Integer.MAX_VALUE : Integer.MIN_VALUE;
            }
            num = num * 10 + s.charAt(i) - '0';
        }
        return num * flag;
    }
}
