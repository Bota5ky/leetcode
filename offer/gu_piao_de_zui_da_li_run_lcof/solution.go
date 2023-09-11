package offer

// 和121相同
// https://leetcode.cn/problems/gu-piao-de-zui-da-li-run-lcof/
func maxProfit(prices []int) int {
	minStack := make([]int, len(prices))
	max := 0
	for i := 0; i < len(prices); i++ {
		if i == 0 || i > 0 && prices[i] < minStack[i-1] {
			minStack[i] = prices[i]
		} else {
			minStack[i] = minStack[i-1]
			if max < prices[i]-minStack[i] {
				max = prices[i] - minStack[i]
			}
		}
	}
	return max
}
