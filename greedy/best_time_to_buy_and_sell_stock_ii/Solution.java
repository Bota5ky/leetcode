package greedy.best_time_to_buy_and_sell_stock_ii;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/">122. 买卖股票的最佳时机 II</a>
 * @since 2023-09-17 14:27
 */
class Solution {
    public int maxProfit(int[] prices) {
        var profit = 0;
        for (int i = 1; i < prices.length; i++) {
            profit += prices[i] > prices[i - 1] ? prices[i] - prices[i - 1] : 0;
        }
        return profit;
    }
}
