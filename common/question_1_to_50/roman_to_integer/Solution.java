package common.question_1_to_50.roman_to_integer;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/roman-to-integer/">13. 罗马数字转整数</a>
 * @since 2023-09-19 21:46
 */
class Solution {
    public int romanToInt(String s) {
        var res = 0;
        var preNum = 1001;
        for (int i = 0; i < s.length(); i++) {
            var num = translate(s.charAt(i));
            if (num > preNum) {
                res -= preNum;
            } else {
                res += preNum;
            }
            preNum = num;
        }
        return res + preNum;
    }

    private int translate(char c) {
        switch (c) {
            case 'I' -> {
                return 1;
            }
            case 'V' -> {
                return 5;
            }
            case 'X' -> {
                return 10;
            }
            case 'L' -> {
                return 50;
            }
            case 'C' -> {
                return 100;
            }
            case 'D' -> {
                return 500;
            }
            default -> {
                return 1000;
            }
        }
    }
}
