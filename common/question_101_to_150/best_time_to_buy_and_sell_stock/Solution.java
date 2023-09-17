package common.question_101_to_150.best_time_to_buy_and_sell_stock;

import java.util.Stack;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/">121. 买卖股票的最佳时机</a>
 * @since 2023-09-17 14:03
 */
class Solution {
    public int maxProfit(int[] prices) {
        var i = 0;
        var profit = 0;
        for (int j = 1; j < prices.length; j++) {
            if (prices[j] <= prices[i]) {
                i = j;
            } else {
                profit = Math.max(profit, prices[j] - prices[i]);
            }
        }
        return profit;
    }
}