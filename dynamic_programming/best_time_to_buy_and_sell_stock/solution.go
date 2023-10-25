package best_time_to_buy_and_sell_stock

// 剑指 Offer 63. 股票的最大利润 https://leetcode.cn/problems/gu-piao-de-zui-da-li-run-lcof/
func maxProfit2(prices []int) int {
	stack1 := make([]int, len(prices))
	stack2 := make([]int, len(prices))
	for i := 0; i < len(prices); i++ {
		if i == 0 || prices[i] < stack1[i-1] {
			stack1[i] = prices[i]
		} else {
			stack1[i] = stack1[i-1]
		}
	}
	for i := len(prices) - 1; i >= 0; i-- {
		if i == len(prices)-1 || prices[i] > stack1[i+1] {
			stack2[i] = prices[i]
		} else {
			stack2[i] = stack2[i+1]
		}
	}
	maximum := 0
	for i := 0; i < len(prices); i++ {
		if stack2[i]-stack1[i] > maximum {
			maximum = stack2[i] - stack1[i]
		}
	}
	return maximum
}
