package bitwise_operation.maximum_odd_binary_number;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/maximum-odd-binary-number/">2864. 最大二进制奇数</a>
 * @since 2024-03-13 18:54
 */
class Solution {
    public String maximumOddBinaryNumber(String s) {
        int ones = 0;
        for (int i = 0; i < s.length(); i++) {
            if (s.charAt(i) == '1') {
                ones++;
            }
        }
        StringBuilder sb = new StringBuilder();
        int zeros = s.length() - ones;
        while (ones > 1) {
            sb.append("1");
            ones--;
        }
        while (zeros > 0) {
            sb.append("0");
            zeros--;
        }
        return sb.append("1").toString();
    }
}
