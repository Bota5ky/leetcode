package common.question_1_to_50.zigzag_conversion;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/zigzag-conversion/">6. N 字形变换</a>
 * @since 2023-09-21 15:31
 */
class Solution {
    public String convert(String s, int numRows) {
        if (numRows == 1) {
            return s;
        }
        var size = numRows * 2 - 2;  //8
        var sb = new StringBuilder();
        for (int row = 0; row < numRows; row++) {
            for (int i = row; i < s.length(); i += size) {
                sb.append(s.charAt(i));
                if (row > 0 && row < size / 2 && i + size - row * 2 < s.length()) {
                    sb.append(s.charAt(i + size - row * 2));
                }
            }
        }
        return sb.toString();
    }
}
