package gu_piao_de_zui_da_li_run_lcof

// 剑指 Offer 63. 股票的最大利润 https://leetcode.cn/problems/gu-piao-de-zui-da-li-run-lcof/
// 121. 买卖股票的最佳时机 https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
func maxProfit(prices []int) int {
	minStack := make([]int, len(prices))
	maximum := 0
	for i := 0; i < len(prices); i++ {
		if i == 0 || i > 0 && prices[i] < minStack[i-1] {
			minStack[i] = prices[i]
		} else {
			minStack[i] = minStack[i-1]
			if maximum < prices[i]-minStack[i] {
				maximum = prices[i] - minStack[i]
			}
		}
	}
	return maximum
}
