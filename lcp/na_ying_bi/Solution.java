package lcp.na_ying_bi;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/na-ying-bi/">LCP 06. 拿硬币</a>
 * @since 2023-09-20 09:14
 */
class Solution {
    public int minCount(int[] coins) {
        var times = 0;
        for (int coin : coins) {
            times += (coin + 1) / 2;
        }
        return times;
    }
}
