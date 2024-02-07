package sort.most_profit_assigning_work;

import java.util.Arrays;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/most-profit-assigning-work/">826. 安排工作以达到最大收益</a>
 * @since 2024-02-07 16:12:01
 */
class Solution {
    public int maxProfitAssignment(int[] difficulty, int[] profit, int[] worker) {
        int[][] dp = new int[difficulty.length][2];
        for (int i = 0; i < dp.length; i++) {
            dp[i] = new int[]{difficulty[i], profit[i]};
        }
        // high profit and low difficulty first
        Arrays.sort(dp, (o1, o2) -> {
            if (o1[1] == o2[1]) {
                return o1[0] - o2[0];
            }
            return o2[1] - o1[1];
        });
        Arrays.sort(worker);
        int sum = 0;
        int i = 0;
        for (int j = worker.length - 1; j >= 0; ) {
            int w = worker[j];
            if (w >= dp[i][0]) {
                sum += dp[i][1];
                j--;
            } else {
                i++;
            }
            if (i >= dp.length) {
                break;
            }
        }
        return sum;
    }
}
