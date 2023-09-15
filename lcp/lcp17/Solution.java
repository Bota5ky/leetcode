package lcp.lcp17;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/nGK0Fy/">LCP 17. 速算机器人</a>
 * @since 2023-09-06 23:23
 */
class Solution {
    public int calculate(String s) {
        int x = 1, y = 0;
        for (int i = 0; i < s.length(); i++) {
            if (s.charAt(i) == 'A') {
                x = 2 * x + y;
            } else {
                y = 2 * y + x;
            }
        }
        return x + y;
    }
}
