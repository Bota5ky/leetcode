package best_time_to_buy_and_sell_stock_ii

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/
func maxProfit(prices []int) int {
	profit := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1]-prices[i] > 0 {
			profit += prices[i+1] - prices[i]
		}
	}
	return profit
}
