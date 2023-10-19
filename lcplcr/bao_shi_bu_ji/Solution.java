package lcplcr.bao_shi_bu_ji;

import java.util.Arrays;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/WHnhjV/">LCP 50. 宝石补给</a>
 * @since 2023-09-15 11:06
 */
class Solution {
    public int giveGem(int[] gem, int[][] operations) {
        for (int[] operation : operations) {
            gem[operation[1]] += gem[operation[0]] / 2;
            gem[operation[0]] -= gem[operation[0]] / 2;
        }
        Arrays.sort(gem);
        return gem[gem.length - 1] - gem[0];
    }
}
