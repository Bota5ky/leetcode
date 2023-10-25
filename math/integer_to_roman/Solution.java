package math.integer_to_roman;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/integer-to-roman/">12. 整数转罗马数字</a>
 * @since 2023-09-19 22:16
 */
class Solution {
    public String intToRoman(int num) {
        var sb = new StringBuilder();
        sb.append("M".repeat(Math.max(0, num / 1000)));
        num %= 1000;
        var hundreds = num / 100;
        // I~III IV VIII IX
        if (hundreds >= 1 && hundreds <= 3) {
            sb.append("C".repeat(hundreds));
        } else if (hundreds == 4) {
            sb.append("CD");
        } else if (hundreds >= 5 && hundreds <= 8) {
            sb.append("D");
            sb.append("C".repeat(hundreds - 5));
        } else if (hundreds == 9) {
            sb.append("CM");
        }
        num %= 100;
        var tens = num / 10;
        // I~III IV VIII IX
        if (tens >= 1 && tens <= 3) {
            sb.append("X".repeat(tens));
        } else if (tens == 4) {
            sb.append("XL");
        } else if (tens >= 5 && tens <= 8) {
            sb.append("L");
            sb.append("X".repeat(tens - 5));
        } else if (tens == 9) {
            sb.append("XC");
        }
        num %= 10;
        if (num >= 1 && num <= 3) {
            sb.append("I".repeat(num));
        } else if (num == 4) {
            sb.append("IV");
        } else if (num >= 5 && num <= 8) {
            sb.append("V");
            sb.append("I".repeat(num - 5));
        } else if (num == 9) {
            sb.append("IX");
        }
        return sb.toString();
    }
}
